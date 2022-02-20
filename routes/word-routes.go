package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterWordStoreRoutes = func(router *gin.Engine, authGroup *gin.RouterGroup) {
	router.GET("/words", controllers.GetAllWords)
	router.GET("/words/:categoryId", controllers.GetAllWordsByCategory)
	router.GET("/word/:wordId", controllers.GetWordById)
	authGroup.POST("/word", controllers.CreateWord)
	authGroup.PUT("/word/:wordId", controllers.UpdateWord)
	authGroup.DELETE("/word/:wordId", controllers.DeleteWord)
}
