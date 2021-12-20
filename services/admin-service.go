package services

import (
	"log"

	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const emailOrPasswordInvalid = "Email or Password is invalid"
const secret = "somuchsecret"

func AdminRegister(
	repository *repositories.AdminRepository,
	admin models.Admin) dtos.Response {
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.PasswordString), 14)

	admin.Password = password
	operationResult := repository.Save(admin)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Admin)

	return dtos.Response{Success: true, Data: data}
}

func AdminLogin(
	repository *repositories.AdminRepository,
	admin models.Admin) dtos.Response {

	response := dtos.Response{Success: true}

	operationResult := repository.GetByEmail(admin.Email)
	if operationResult.Error != nil {
		response = dtos.Response{Success: false, Message: operationResult.Error.Error()}
		return response
	}

	var existingAdmin = operationResult.Result.(*models.Admin)
	if existingAdmin.ID == 0 {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return response
	}

	log.Println(existingAdmin)
	if err := bcrypt.CompareHashAndPassword(existingAdmin.Password, []byte(admin.PasswordString)); err != nil {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return response
	}

	return response
}

func AdminLogout() {

}

func IsAuthentificated(cookie string) (bool, *jwt.Token) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, token
	}

	return true, token
}
