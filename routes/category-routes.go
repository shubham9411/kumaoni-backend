package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterCategoryStoreRoutes = func(router *gin.Engine) {
	router.GET("/categories", controllers.GetAllCategories)
}
