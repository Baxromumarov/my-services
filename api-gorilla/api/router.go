package api

import (
	"log"
	"net/http"

	v1 "github.com/baxromumarov/my-services/api-gateway/api/handlers/v1"
	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/services"

	
	"github.com/gorilla/mux"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	router.HandleFunc("/v1/users/create",handlerV1.CreateUser).Methods("POST")
	router.HandleFunc("/v1/users/get",handlerV1.GetUser).Methods("GET")
	// router.HandleFunc("/v1/users/delete",handlerV1.DeleteUser).Methods("Delete")
	// router.HandleFunc("/v1/users/update",handlerV1.UpdateUser).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(config.Load().HTTPPort, router))
	
	return router
}
