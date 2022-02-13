package controllers

import (
	"log"
	"net/http"

	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKeyF = "id"

type ForemanController struct {
	foremanRepository *repositories.ForemanRepository
}

func NewForemanController(foremanRepository *repositories.ForemanRepository) *ForemanController {
	return &ForemanController{foremanRepository: foremanRepository}
}

func (c *ForemanController) RegisterForemanHandler(context *gin.Context) {
	var foreman models.Foreman

	err := context.ShouldBindJSON(&foreman)

	// validation errors
	if err != nil {
		log.Println("Cannot unmarshal admin data, error: ", err.Error())
		// generate validation errors response
		response := helpers.GenerateValidationResponse(err)

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	// save project & get it's response
	response := services.ForemanRegister(c.foremanRepository, foreman)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ForemanController) LoginForemanHandler(loginVals *models.LoginInfo) dtos.Response {
	return services.ForemanLogin(c.foremanRepository, loginVals)
}

func (c *ForemanController) GetForemanInfoHandler(context *gin.Context) {
	code := http.StatusOK
	claims := jwt.ExtractClaims(context)
	response := &dtos.Response{
		Success: true,
		Data: &models.LoginInfo{
			Email: claims[identityKeyF].(string),
		},
	}

	context.JSON(code, response)
}
