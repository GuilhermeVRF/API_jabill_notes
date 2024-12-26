package controllers

import (
	"jabill-notes/auth"
	"jabill-notes/services"
	"jabill-notes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaController struct{
	mediaService services.MediaService
}

func NewMediaController(mediaService services.MediaService) MediaController{
	return MediaController{
		mediaService: mediaService,
	}
}

func (mediaController *MediaController) Show(context *gin.Context){
	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	userProfile, err := mediaController.mediaService.GetUserProfile(user.Id);
	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))	
	} 

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "User profile retrieved!", gin.H{ "profile": userProfile}))
}