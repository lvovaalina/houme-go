package main

import (
	"log"
	"os"

	"./forge"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

var Projects []forge.Project

func getAllProjects(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	Projects = forge.GetBucketObjects("houme")
	c.JSON(200, Projects)
}

func main() {
	port := os.Getenv("PORT")
	log.Println("dfd")
	if port == "" {
		port = "10000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/projects", getAllProjects)

	router.Run(":" + port)
}
