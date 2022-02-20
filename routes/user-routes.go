package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/controllers"
	"github.com/shubham9411/kumaoni-backend/middlewares"
)

var RegisterUserRoutes = func(router *gin.Engine, authGroup *gin.RouterGroup) {
	authGroup.GET("/users/", controllers.GetAllUsers, middlewares.Authenticate)
	authGroup.GET("/users/:user_id", controllers.GetUser, middlewares.Authenticate)
}
