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

var identityKey = "id"

type AdminController struct {
	adminRepository *repositories.AdminRepository
}

func NewAdminController(adminRepository *repositories.AdminRepository) *AdminController {
	return &AdminController{adminRepository: adminRepository}
}

func (c *AdminController) RegisterAdminHandler(context *gin.Context) {
	var admin models.Admin

	err := context.ShouldBindJSON(&admin)

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
	response := services.AdminRegister(c.adminRepository, admin)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *AdminController) LoginAdminHandler(loginVals models.Admin) dtos.Response {
	return services.AdminLogin(c.adminRepository, loginVals)
}

func (c *AdminController) GetAdminInfoHandler(context *gin.Context) {
	code := http.StatusOK
	claims := jwt.ExtractClaims(context)
	response := &dtos.Response{
		Success: true,
		Data: &models.Admin{
			Email: claims[identityKey].(string),
		},
	}

	context.JSON(code, response)
}
