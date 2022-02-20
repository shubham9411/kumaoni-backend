package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterPhraseStoreRoutes = func(router *gin.Engine, authGroup *gin.RouterGroup) {
	router.GET("/phrases", controllers.GetAllPhrases)
	router.GET("/phrases/:categoryId", controllers.GetAllPhrasesByCategory)
	router.GET("/phrase/:phraseId", controllers.GetPhraseById)
	authGroup.POST("/phrase", controllers.CreatePhrase)
	authGroup.PUT("/phrase/:phraseId", controllers.UpdatePhrase)
	authGroup.DELETE("/phrase/:phraseId", controllers.DeletePhrase)
}
