package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
)

var RegisterAuthRoutes = func(router *gin.Engine, authGroup *gin.RouterGroup) {
	router.POST("/users/signup", controllers.Signup)
	router.POST("/users/login", controllers.Login)
}
