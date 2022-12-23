package main

import (
	"fmt"

	"github.com/baxromumarov/my-services/api-gateway/api"
	handle "github.com/baxromumarov/my-services/api-gateway/api/handlers"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	"github.com/gin-gonic/gin"
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

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	// cfg := config.Load()
	// log := logger.New(cfg.LogLevel, "My_Api_Gateway")

	// pool := redis.Pool{

	// 	MaxIdle:   80,

	// 	MaxActive: 12000,

	// 	Dial: func() (redis.Conn, error) {
	// 		c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		return c, err
	// 	},
	// }

	// redisRepo := rds.NewRedisRepo(&pool)

	// serviceManager, err := services.NewServiceManager(&cfg)

	// if err != nil {
	// 	log.Error("gRPC dial error", logger.Error(err))
	// }

	// server := api.SetUpApi(api.Option{
	// 	Conf:           cfg,
	// 	Logger:         log,
	// 	ServiceManager: serviceManager,
	// 	RedisRepo: redisRepo,
	// })

	// if err := server.Run(cfg.HTTPPort); err != nil {
	// 	log.Fatal("failed to run http server", logger.Error(err))
	// 	panic(err)
	// }
	cfg := config.Load()

	grpcSvcs, err := services.NewServiceManager(&cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("tu_go_admin_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handle.NewHandlerV1(&cfg, log, grpcSvcs)

	api.SetUpApi(&cfg, r, *h)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
