package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"bitbucket.org/houmeteam/houme-go/forge"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type TranslateResponse struct {
	FileName string `json:"filename"`
}

func getAllProjects(c *gin.Context) {
	projects := forge.GetBucketObjects("houme")
	c.JSON(http.StatusOK, projects)
}

func uploadFile(c *gin.Context) {
	log.Println("Starting to upload a file")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	log.Println(filename)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Could not copy bytes. File err : %s", err.Error()))
		return
	}

	forge.UploadFileBinaryToBucket("houme", buf.Bytes(), filename)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded file"})
}

func translateFile(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
	}

	var result TranslateResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	forge.TranslateFile("houme", result.FileName)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func deleteFile(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
	}

	var result TranslateResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	forge.DeleteFileInBucket("houme", result.FileName)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/projects", getAllProjects)
	router.POST("/upload", uploadFile)
	router.POST("/translate", translateFile)
	router.DELETE("/deleteFile", deleteFile)

	router.Run(":" + port)
}
