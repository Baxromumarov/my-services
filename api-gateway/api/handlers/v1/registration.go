package v1

import (
	"context"
	"math/rand"
	"strconv"
	"strings"

	"net/http"
	"time"

	pb "github.com/baxromumarov/my-services/api-gateway/genproto"
	l "github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"google.golang.org/protobuf/encoding/protojson"

	"crypto/tls"
	"fmt"

	// "github.com/google/uuid"
	gomail "gopkg.in/mail.v2"
)

// @Summary Register User
// @Description This api uses for registration new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequestBody true "user body"
// @Success 200 {string} Success
// @Router /v1/users/register [post]
func (h handlerV1) RegisterUser(c *gin.Context) {
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

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//----------------------------------------------------------------
	status, err := h.serviceManager.UserService().CheckField(ctx, &pb.UserCheckRequest{
		Field: "username",
		Value: body.UserName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while calling CheckField function whith USERNAME", l.Error(err))
		return
	}

	if !status.Response {
		status2, err := h.serviceManager.UserService().CheckField(ctx, &pb.UserCheckRequest{
			Field: "email",
			Value: body.Email,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while calling CheckField function with EMAIL", l.Error(err))
			return
		}

		if status2.Response {
			c.JSON(http.StatusConflict, gin.H{
				"error": "user_name already in use",
			})
			h.log.Error("User already exists", l.Error(err))
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email already in use",
		})
		h.log.Error("User already exists", l.Error(err))
		return
	}

	min := 99999
	max := 1000000
	rand.Seed(time.Now().UnixNano())
	Code := rand.Intn(max-min) + min

	verCode := strconv.Itoa(Code)

	body.EmailCode = verCode

	SendEmail(body.Email, verCode)

	setBodyRedis, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed set to redis IN REGISTER FUNC  1", l.Error(err))
		return
	}

	err = h.redisStorage.Set(body.Email, string(setBodyRedis))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed set to redis IN REGISTER FUNC  2", l.Error(err))
		return
	}

}

// @Summary Send Email Code
// @Description This api uses for sendin email code to user
// @Tags users
// @Accept json
// @Produce json
// @Param user body EmailVer true "user body"
// @Success 200 {string} Success
// @Router /v1/users/verification [post]
func (h handlerV1) VerifyUser(c *gin.Context) {

	var mailData EmailVer

	err := c.ShouldBindJSON(&mailData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json in VerifyUser func", l.Error(err))
		return
	}


	mailData.Email = strings.TrimSpace(mailData.Email)
	mailData.Email = strings.ToLower(mailData.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	bod, err := redis.String(h.redisStorage.Get(mailData.Email))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		h.log.Error("failed get from redis", l.Error(err))
		return
	}

	var redisBody *pb.User

	err = json.Unmarshal([]byte(bod), &redisBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while using CreateUser func in verify", l.Error(err))
		return
	}

	// valid:= ValidPassword(redisBody.EmailCode) 
	
	if mailData.EmailCode == redisBody.EmailCode  {
		createVal, err := h.serviceManager.UserService().Insert(ctx, redisBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("error while inserting db", l.Error(err))
			return
		}

		c.JSON(http.StatusCreated, createVal)
	}
}

func ValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	  }
	  //check password contain lowercase letter
	  if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	  }
	  //check password contain uppercase letter
	  if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	  }
	  //check password contain number
	  if !strings.ContainsAny(password, "0123456789") {
		return false
	  }
	  //check password contain special character
	  if !strings.ContainsAny(password, "!@#$%^&*()_+") {
		return false
	  }
	  return true
}

func SendEmail(email, code string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "ahrorahrorovnt@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)
	// id,err := uuid.NewUUID()
	// if err != nil {
	//   fmt.Println(err)
	// }
	// Set E-Mail subject
	m.SetHeader("code:", "Verification code")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", code)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "ahrorahrorovnt@gmail.com", "qmxlgijkvuuoacrh")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
