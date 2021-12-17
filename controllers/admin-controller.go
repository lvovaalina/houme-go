package controllers

import (
	"log"
	"net/http"

	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	"github.com/gin-gonic/gin"
)

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

func (c *AdminController) LoginAdminHandler(context *gin.Context) {
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
	response, token := services.AdminLogin(c.adminRepository, admin)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.SetSameSite(http.SameSiteNoneMode)

	context.SetCookie("jwt", token, 60*60*2, "/", "https://houmly-dev.herokuapp.com", true, true)
	context.Writer.Header().Add("access-control-expose-headers", "Set-Cookie")
	context.JSON(code, response)
}
