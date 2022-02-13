package services

import (
	"log"
	"strings"

	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func ForemanRegister(
	repository *repositories.ForemanRepository,
	Foreman models.Foreman) dtos.Response {
	password, _ := bcrypt.GenerateFromPassword([]byte(Foreman.PasswordString), 14)

	Foreman.Password = password
	operationResult := repository.Save(Foreman)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Foreman)

	return dtos.Response{Success: true, Data: data}
}

func ForemanLogin(
	repository *repositories.ForemanRepository,
	foreman *models.LoginInfo) dtos.Response {

	response := dtos.Response{Success: true}

	operationResult := repository.GetByEmail(strings.ToLower(foreman.Email))
	if operationResult.Error != nil {
		response = dtos.Response{Success: false, Message: operationResult.Error.Error()}
		return response
	}

	var existingForeman = operationResult.Result.(*models.Foreman)
	if existingForeman.ID == 0 {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return response
	}

	log.Println(existingForeman)
	if err := bcrypt.CompareHashAndPassword(existingForeman.Password, []byte(foreman.PasswordString)); err != nil {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return response
	}

	return response
}

func ForemanLogout() {

}

func IsForemanAuthentificated(cookie string) (bool, *jwt.Token) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, token
	}

	return true, token
}
