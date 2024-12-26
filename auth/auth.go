package auth

import (
	"errors"
	"jabill-notes/models"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func ParseToken(authorization string) (models.User, error){
	tokenString, err := getBearerToken(authorization)

	if err != nil{
		return models.User{}, err
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("Invalid JWT signature method!")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil{
		return models.User{}, err
	}

	if userClaims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid{
		user := models.User{
			Id: int(userClaims["id"].(float64)),
			Email: userClaims["email"].(string),
			Name:  userClaims["name"].(string),
		}

		return user, nil
	}

	return models.User{}, errors.New("Invalid user token!")
}

func getBearerToken(authorization string) (string, error){
	if authorization == ""{
		return "", errors.New("Bearer Token ausente na requisição!")
	}
	
	splittedToken := strings.Split(authorization, " ")

	if len(splittedToken) != 2 || splittedToken[0] != "Bearer"{
		return "", errors.New("Formato de authorization incorreto!")	
	}

	return splittedToken[1], nil
}