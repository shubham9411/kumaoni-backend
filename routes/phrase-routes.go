package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterPhraseStoreRoutes = func(router *gin.Engine) {
	router.POST("/phrase", controllers.CreatePhrase)
	router.GET("/phrases", controllers.GetAllPhrases)
	router.GET("/phrases/:categoryId", controllers.GetAllPhrasesByCategory)
	router.GET("/phrase/:phraseId", controllers.GetPhraseById)
	router.PUT("/phrase/:phraseId", controllers.UpdatePhrase)
	router.DELETE("/phrase/:phraseId", controllers.DeletePhrase)
}
