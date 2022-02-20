package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
	"github.com/shubham9411/kumaoni-backend/utils"
)

func GetAllCategories(c *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, categories)
}
