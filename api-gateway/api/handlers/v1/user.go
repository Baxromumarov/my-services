package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/baxromumarov/my-services/api-gateway/genproto"
	l "github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/baxromumarov/my-services/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type EmailVer struct {
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email"`
	EmailCode string `protobuf:"bytes,15,opt,name=email_code,json=emailCode,proto3" json:"email_code"`
}

type CreateUserRequestBody struct {
	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName    string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName     string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	Email        string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email"`
	Bio          string     `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio"`
	PhoneNumbers []string   `protobuf:"bytes,6,rep,name=phoneNumbers,proto3" json:"phoneNumbers"`
	Addresses    []*Address `protobuf:"bytes,7,rep,name=Addresses,proto3" json:"Addresses"`
	Post         []*Post    `protobuf:"bytes,8,rep,name=post,proto3" json:"post"`
	TypeId       int64      `protobuf:"varint,9,opt,name=typeId,proto3" json:"typeId"`
	Status       string     `protobuf:"bytes,10,opt,name=Status,proto3" json:"Status"`
	CreatedAt    string     `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt    string     `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt"`
	DeletedAt    string     `protobuf:"bytes,13,opt,name=deletedAt,proto3" json:"deletedAt"`
	UserName     string    `protobuf:"bytes,13,opt,name=user_name,json=userName,proto3" json:"user_name"`
}

type Address struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	City       string `protobuf:"bytes,2,opt,name=City,proto3" json:"City"`
	Country    string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country"`
	District   string `protobuf:"bytes,4,opt,name=District,proto3" json:"District"`
	PostalCode string `protobuf:"bytes,5,opt,name=PostalCode,proto3" json:"PostalCode"`
}

type Post struct {
	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Medias    []*Media `protobuf:"bytes,4,rep,name=medias,proto3" json:"medias"`
	CreatedAt string   `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt"`
}

type Media struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Link string `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
}

// CreateUser creates user
// @Summary Create user
// @Description This api is using for creating new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body CreateUserRequestBody true "user body"
// @Success 200 {string} Success
// @Router /v1/users/post [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	fmt.Println(&body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Insert(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by id
// @Summary Get user
// @Description This api is using for getting user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {string} CreateUserRequestBody
// @Router /v1/users/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetById(
		ctx, &pb.ById{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetUserList gets users list
// @Summary Get user list
// @Description This api is using for getting users list
// @Tags user
// api.DELETE("/users/:id", handlerV1.DeleteUser)
// @Accept  json
// @Produce  json
// @Param limit path int true "limit"
// @Param page path int true "page"
// @Success 200 {string} CreateUserRequestBody
// @Router /v1/users [get]
func (h *handlerV1) ListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().ListUsers(
		ctx, &pb.GetUsersRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// // UpdateUser updates user by id
// @Summary Update user
// @Description This api is using for updatting user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {string} CreateUserRequestBody
// @Router /v1/users/update/{id} [get]
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
