package v1

import (
	"errors"
	"net/http"

	"github.com/baxromumarov/my-services/api-gateway/api/auth"
	model "github.com/baxromumarov/my-services/api-gateway/api/model"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	repo "github.com/baxromumarov/my-services/api-gateway/storage/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redisStorage   repo.RedisRepositoryStorage
	jwtHandler auth.JwtHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.RedisRepositoryStorage
	jwtHandler auth.JwtHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage:   c.Redis,
		jwtHandler: c.jwtHandler,
	}
}

func CheckClaims(h *handlerV1, c *gin.Context) jwt.MapClaims{
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization model.JwtRequestModel
		claims jwt.MapClaims
		err error
	)
	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == ""{
		c.JSON(http.StatusUnauthorized,ErrUnauthorized)
		h.log.Error("Unauthorized request",logger.Error(ErrUnauthorized))
		return nil
	}
	h.jwtHandler.Token = authorization.Token
	claims, err = h.jwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized,ErrUnauthorized)
		h.log.Error("Unauthorized request",logger.Error(ErrUnauthorized))
		return nil
	}
	return claims
}

