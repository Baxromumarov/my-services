package main

import (
	"fmt"
	"github.com/baxromumarov/my-services/api-gateway/api"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	rds "github.com/baxromumarov/my-services/api-gateway/storage/redis"
	casbin"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gomodule/redigo/redis"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8010
// @BasePath  /api/v1

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "My_Api_Gateway")

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := gormadapter.NewAdapter("postgres", psqlString, true)

	if err != nil {
		log.Error("new adapter error", logger.Error(err))
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, db)
	if err != nil {
		log.Error("new enforcer error", logger.Error(err))
		return
	}
	/*
	casbinEnforcer,err = casbin.NewEnforcer(cfg.CasbinConfigPathCsv,"config/rbac_policy.csv")
	if err != nil {
		log.Error("new enforcer error", logger.Error(err))
		return
	}
	*/
	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("new load policy error", logger.Error(err))
		return
	}

	pool := redis.Pool{

		MaxIdle: 80,

		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	redisRepo := rds.NewRedisRepo(&pool)

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch3", util.KeyMatch3)

	serviceManager, err := services.NewServiceManager(&cfg)

	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Casbin:         casbinEnforcer,
		ServiceManager: serviceManager,
		RedisRepo:      redisRepo,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
