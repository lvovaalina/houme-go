package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	propertiesRepository *repositories.PropertyRepository
	jobsRepository       *repositories.JobRepository
}

type TranslateResponse struct {
	FileName string `json:"filename"`
}

func NewCommonController(propertiesRepository *repositories.PropertyRepository,
	jobsRepository *repositories.JobRepository) *CommonController {
	return &CommonController{
		propertiesRepository: propertiesRepository,
		jobsRepository:       jobsRepository,
	}
}

func (c *CommonController) ForgeGetHandler(context *gin.Context) {
	projects := forge.GetBucketObjects("houmly")
	context.JSON(http.StatusOK, projects)
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

	forge.TranslateFile("houme", result.FileName)
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
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
