package api

import (
	repo "github.com/baxromumarov/my-services/api-gateway/storage/repo"

	"github.com/baxromumarov/my-services/api-gateway/api/auth"
	_ "github.com/baxromumarov/my-services/api-gateway/api/docs"
	v1 "github.com/baxromumarov/my-services/api-gateway/api/handlers/v1"
	"github.com/baxromumarov/my-services/api-gateway/api/middleware"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	casbin "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo      repo.RedisRepositoryStorage
	Casbin         *casbin.Enforcer
}

// New @BasePath /v1
// New ...
// @SecurityDefinitions.apikey BearerAuth
// @Description GetMyProfile
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandler := &auth.JwtHandler{
		SigningKey: option.Conf.SigningKey,
		Log:        option.Logger,
	}

	router.Use(middleware.NewJwtRoleStuct(option.Casbin, option.Conf, *jwtHandler))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.RedisRepo,
	})

	api := router.Group("/v1")
	api.POST("/users/verification", handlerV1.VerifyUser)
	api.POST("/users/register", handlerV1.RegisterUser)
	api.GET("/users/login", handlerV1.LoginUser)

	api.POST("/users/post", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.ListUsers)
	api.GET("/v1/users/idtoken", handlerV1.GetUserByIdFromToken)
	// api.PUT("/users/update/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
