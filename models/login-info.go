package models

type LoginInfo struct {
	Email          string
	PasswordString string
	Password       []byte
	Role           string
}
