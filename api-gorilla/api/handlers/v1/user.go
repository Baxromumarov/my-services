package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	pb "github.com/baxromumarov/my-services/api-gateway/genproto"
)

// CreateUser creates user
// route /v1/users [post]
func (h *handlerV1) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		user        pb.User
		
	)
	requestBody, err := ioutil.ReadAll(r.Body) 
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	json.Unmarshal(requestBody, &user)
	w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Insert(
		ctx, 
		&user,
	)
	if err != nil {
		fmt.Fprintf(w, "Error grpc: %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(response)

}

// GetUser gets user by id
// route /v1/users/{id} [get]
func (h *handlerV1) GetUser(w http.ResponseWriter, r *http.Request) {
	var user pb.User
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	json.Unmarshal(requestBody, &user)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetById(
		ctx,
		&pb.ById{
			Id: user.Id,
		},
	)
	if err != nil {
		fmt.Fprintf(w, "Error grpc: %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(response)
}

// // ListUsers returns list of users
// // route /v1/users/ [get]
// func (h *handlerV1) ListUsers(w http.ResponseWriter, r *http.Request) {
// 	queryParams := c.Request.URL.Query()

// 	params, errStr := utils.ParseQueryParams(queryParams)
// 	if errStr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": errStr[0],
// 		})
// 		h.log.Error("failed to parse query params json" + errStr[0])
// 		return
// 	}

// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().List(
// 		ctx, &pb.ListReq{
// 			Limit: params.Limit,
// 			Page:  params.Page,
// 		})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to list users", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // UpdateUser updates user by id
// // route /v1/users/{id} [put]
// func (h *handlerV1) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		body        pb.User
// 		jspbMarshal protojson.MarshalOptions
// 	)
// 	jspbMarshal.UseProtoNames = true

// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}
// 	body.Id = c.Param("id")

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().Update(ctx, &body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to update user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // DeleteUser deletes user by id
// // route /v1/users/{id} [delete]
// func (h *handlerV1) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	guid := c.Param("id")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().Delete(
// 		ctx, &pb.ByIdReq{
// 			Id: guid,
// 		})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to delete user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }
