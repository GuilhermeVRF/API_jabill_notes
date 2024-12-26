package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Profile_image string `json:"profile_image"`
}

func (user *User) GenerateToken() (string, error){
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	userClaims := jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"name": user.Name,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	stringToken, err := userToken.SignedString([]byte(jwtSecret))
	if err != nil{
		return "", err
	}

	return stringToken, nil
}