package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterWordStoreRoutes = func(router *gin.Engine) {
	router.POST("/word", controllers.CreateWord)
	router.GET("/words", controllers.GetAllWords)
	router.GET("/words/:categoryId", controllers.GetAllWordsByCategory)
	router.GET("/word/:wordId", controllers.GetWordById)
	router.PUT("/word/:wordId", controllers.UpdateWord)
	router.DELETE("/word/:wordId", controllers.DeleteWord)
}
