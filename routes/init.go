package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/middlewares"
)

func InitializeRoutes(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(middlewares.Authenticate)

	RegisterAuthRoutes(router, authorized)
	RegisterUserRoutes(router, authorized)
	RegisterWordStoreRoutes(router, authorized)
	RegisterCategoryStoreRoutes(router, authorized)
	RegisterPhraseStoreRoutes(router, authorized)
}
