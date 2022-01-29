package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/routes"
)

func main() {
	fmt.Println("Welcome to Kumaoni API")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterBookStoreRoutes(r)

	err := r.Run(":3000")
	if err != nil {
		log.Fatal("Something went wrong:: ", err)
	}
}
