package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/config"
	"github.com/shubham9411/kumaoni-backend/models"
	"github.com/shubham9411/kumaoni-backend/routes"
	"github.com/shubham9411/kumaoni-backend/utils"
)

func main() {
	fmt.Println("Welcome to Kumaoni API")

	db := config.Connect()
	models.SetDatabase(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.InitializeRoutes(r)

	port := utils.GodotEnv("GO_PORT")

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Something went wrong:: ", err)
	}
}
