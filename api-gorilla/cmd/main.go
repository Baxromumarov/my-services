package main

import (
	"net/http"

	"github.com/baxromumarov/my-services/api-gateway/api"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"
	
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := http.ListenAndServe(cfg.HTTPPort,server); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
