package api

import (
	handle "github.com/baxromumarov/my-services/api-gateway/api/handlers"
	"github.com/baxromumarov/my-services/api-gateway/storage"
	repo "github.com/baxromumarov/my-services/api-gateway/storage/repo"

	_ "github.com/baxromumarov/my-services/api-gateway/api/docs"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.LoggerI
	ServiceManager services.IServiceManager
	RedisRepo      repo.RedisRepositoryStorage
	Storage        storage.IStorage
}

func SetUpApi(cfg *config.Config, r *gin.Engine, h handle.HandlerV1) {


	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")

	v1.POST("/users/verification", h.RegisterUser)
	v1.POST("/users/register", h.RegisterUser)
	v1.POST("/users/post", h.CreateUser)
	v1.GET("/users/:id", h.GetUser)
	v1.GET("/users", h.ListUsers)
	// api.PUT("/users/update/:id", h.UpdateUser)
	// api.DELETE("/users/:id", h.DeleteUser)
	
	v1.GET("/get/json-resp", h.GetJson)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
