package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/baxromumarov/api-gateway/genproto"
	l "github.com/baxromumarov/api-gateway/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
)

//CreateUser creates user
// route /v1/users [post]

func (h *handlerV1) CreateUser(c *gin.Context){
	var (
		body pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	fmt.Println(&body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)

}

// GetUser gets users by id
// route /v1/users/:id [get]

func (h *handlerV1) GetUser(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response,err:= h.serviceManager.UserService().GetUser(
		ctx, 
		&pb.GetUserRequest{
			Id: guid,
		})
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to get user", l.Error(err))
		return
		}
	
	c.JSON(http.StatusOK, response)
}


// // ListUsers returns list of users
// // route /v1/users/ [get]
// func (h *handlerV1) ListUsers(c *gin.Context) {
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
// func (h *handlerV1) UpdateUser(c *gin.Context) {
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
// func (h *handlerV1) DeleteUser(c *gin.Context) {
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