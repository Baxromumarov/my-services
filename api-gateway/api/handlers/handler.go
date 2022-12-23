package v1

import (
	"github.com/baxromumarov/my-services/api-gateway/api/http"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	"github.com/baxromumarov/my-services/api-gateway/storage"

	repo "github.com/baxromumarov/my-services/api-gateway/storage/repo"
	"github.com/gin-gonic/gin"
)

type HandlerV1 struct {
	cfg            *config.Config
	storage        storage.IStorage
	serviceManager services.IServiceManager
	redisStorage   repo.RedisRepositoryStorage
	log            logger.LoggerI
}

func NewHandlerV1(cfg *config.Config, log logger.LoggerI, svcs services.IServiceManager) *HandlerV1 {
	return &HandlerV1{
		cfg:      cfg,
		log:      log,
		serviceManager: svcs,
	}
}

func (h *HandlerV1) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, data)
}
