package api

import (
	v1 "github.com/baxromumarov/my-services/api-gateway/api/handlers/v1"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	"github.com/baxromumarov/my-services/user-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct{
	Conf config.Config
	Logger logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine{
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger: option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg: option.Conf,
	})
	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	// api.GET("/users", handlerV1.ListUsers)
	// api.PUT("/users/:id", handlerV1.UpdateUser)
	// api.DELETE("/users/:id", handlerV1.DeleteUser)

	return router
}