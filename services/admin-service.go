package services

import (
	"strconv"
	"time"

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
	admin models.Admin) (response dtos.Response, token string) {
	token = ""
	response = dtos.Response{Success: true}
	operationResult := repository.GetByEmail(admin.Email)
	if operationResult.Error != nil {
		response = dtos.Response{Success: false, Message: operationResult.Error.Error()}
		return
	}

	var existingAdmin = operationResult.Result.(*models.Admin)
	if existingAdmin.ID == 0 {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return
	}

	if err := bcrypt.CompareHashAndPassword(existingAdmin.Password, []byte(admin.PasswordString)); err != nil {
		response = dtos.Response{Success: false, Message: emailOrPasswordInvalid}
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), //2 hours
	})

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		response = dtos.Response{Success: false, Message: "Could not login"}
	}

	return
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
