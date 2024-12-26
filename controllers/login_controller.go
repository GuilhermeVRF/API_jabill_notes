package controllers

import (
	"jabill-notes/models"
	"jabill-notes/services"
	"jabill-notes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{
	loginService services.LoginService
}

func NewLoginController(loginService services.LoginService) LoginController{
	return LoginController{
		loginService: loginService,
	}
}

func (loginController *LoginController) Verify(context *gin.Context){
	var login models.Login
	err := context.ShouldBindBodyWithJSON(&login)

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), gin.H{"user_token" : nil}))
		return
	}

	userToken, err := loginController.loginService.Verify(login)

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), gin.H{"user_token" : nil}))
		return
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Login realizado com sucesso!", gin.H{"user_token" : userToken}))
}