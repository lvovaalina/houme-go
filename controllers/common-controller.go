package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	propertiesRepository *repositories.PropertyRepository
	jobsRepository       *repositories.JobRepository
	uploader             *helpers.ClientUploader
}

type TranslateResponse struct {
	FileName string `json:"filename"`
}

func NewCommonController(propertiesRepository *repositories.PropertyRepository,
	jobsRepository *repositories.JobRepository,
	clientUploader *helpers.ClientUploader) *CommonController {
	return &CommonController{
		propertiesRepository: propertiesRepository,
		jobsRepository:       jobsRepository,
		uploader:             clientUploader,
	}
}

func (c *CommonController) ForgeGetHandler(context *gin.Context) {
	projects := forge.GetBucketObjects("houmly")
	context.JSON(http.StatusOK, projects)
}

func (c *CommonController) UploadModel(context *gin.Context) {
	name := context.PostForm("name")
	email := context.PostForm("email")

	f, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Strarting to upload file: " + f.Filename)

	filename := time.Now().String() + f.Filename

	blobFile, err := f.Open()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.uploader.UploadFile(blobFile, filename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.uploader.CreateFile(filename, name, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "success",
	})
}

func (c *CommonController) ForgeUploadHandler(context *gin.Context) {
	log.Println("Starting to upload a file")
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	log.Println(filename)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("Could not copy bytes. File err : %s", err.Error()))
		return
	}

	forge.UploadFileBinaryToBucket("houmly", buf.Bytes(), filename)
	context.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded file"})
}

func (c *CommonController) ForgeTranslateHandler(context *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
	}

	var result TranslateResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	forge.TranslateFile("houmly", result.FileName)
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (c *CommonController) ForgeTranslationStatusHandler(context *gin.Context) {

	forge.GetTranslationStatus("houmly", "Classic_house.rvt")
	context.JSON(http.StatusOK, gin.H{
		"message": "ok"})
}

func (c *CommonController) ForgeDeleteFile(context *gin.Context) {
	filename := context.Param("filename")
	log.Println("FROM REQUEST", filename)

	forge.DeleteFileInBucket("houmly", filename)
	context.JSON(http.StatusOK, gin.H{
		"message": "ok"})
}

func (c *CommonController) GetProperties(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllProperties(c.propertiesRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CommonController) GetJobsHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllJobs(c.jobsRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
