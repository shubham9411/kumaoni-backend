package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	// Method     string      `json:"method"`
}

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	// Method     string      `json:"method"`
}

func SendError(e string, c *gin.Context) {
	fmt.Println(e)
	err := ErrorResponse{StatusCode: http.StatusUnprocessableEntity, Error: e}
	c.JSON(http.StatusUnprocessableEntity, err)
}
