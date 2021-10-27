package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/model"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

type TranslateResponse struct {
	FileName string `json:"filename"`
}

type ProjectModel struct {
	Name                     string `json:"name"`
	LivingArea               string `json:"livingArea"`
	RoomsNumber              int    `json:"roomsNumber"`
	ConstructonWorkersNumber string `json:"constructonWorkersNumber"`
	FoundationMaterial       string `json:"foundationMaterial"`
	WallMaterial             string `json:"wallMaterial"`
	FinishMaterial           string `json:"finishMaterial"`
	RoofingMaterial          string `json:"roofingMaterial"`
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

func addProject(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
	}

	var result ProjectModel
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	project := model.Project{
		Name:                     result.Name,
		LivingArea:               result.LivingArea,
		RoomsNumber:              result.RoomsNumber,
		ConstructonWorkersNumber: result.ConstructonWorkersNumber,
		WallMaterial:             result.WallMaterial,
		FinishMaterial:           result.FinishMaterial,
		FoundationMaterial:       result.FoundationMaterial,
		RoofingMaterial:          result.RoofingMaterial,
	}

	dbResult := db.Create(&project)
	if dbResult.Error == nil {
		c.JSON(http.StatusOK, gin.H{"project": project})
	} else {
		log.Println(dbResult.Error, project.ID)
	}
}

func getProperties(c *gin.Context) {

}

func RouterInitialize(port string) {
	router := gin.New()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=houmly sslmode=disable password=l8397040")

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.ProjectProperty{})
	db.AutoMigrate(&model.Property{})

	//handler := cors.Default().Handler(router)
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
	router.POST("/add", addProject)
	router.GET("/properties", getProperties)

	router.Run(":" + port)
}
