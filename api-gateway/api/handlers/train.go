package v1

import (
	"fmt"
	"net/http"

	"github.com/baxromumarov/my-services/api-gateway/models"
	"github.com/gin-gonic/gin"
)

// GetJson getdata
// @Router /v1/get/json-resp [GET]
// @Summary get json data
// @Description API for getting json data
// @Tags train
// @Accept json
// @Produce json
// @Success 200 {object} models.JsonFile
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *HandlerV1) GetJson(c *gin.Context) {
	var r models.JsonFile
	fmt.Println(r)

	h.storage.Train().GetJson()
	// if err != nil {
	// 	fmt.Println("error in api", err)
	// 	return
	// }
	// fmt.Println(resp)
	c.JSON(http.StatusOK, "resp")
}
