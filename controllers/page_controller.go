package controllers

import (
	"jabill-notes/auth"
	"jabill-notes/models"
	"jabill-notes/requests"
	"jabill-notes/services"
	"jabill-notes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageController struct{
	pageService services.PageService
}

func NewPageController (pageService services.PageService) PageController{
	return PageController{
		pageService: pageService,
	}
}

func (pageController *PageController) Show(context *gin.Context){
	authorizationHeader := context.GetHeader("Authorization")
	slug := context.Param("slug")

	if slug == ""{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "Parâmetro slug não encontrado!", nil))
		return 	
	}

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	page, err := pageController.pageService.Show(slug, user.Id)

	if err != nil{
		context.JSON(http.StatusNotFound, utils.NewApiResponse("error", err.Error(), nil))
		return
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Página coletada com sucesso!", page))
}

func (pageController *PageController) Index(context *gin.Context){
	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	pages, err := pageController.pageService.Index(user.Id)

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), []interface{}{}))	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Páginas do usuário coletados com sucesso!", pages))
}


func (pageController *PageController) Store(context *gin.Context){
	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	var page models.Page
	
	err = context.ShouldBindJSON(&page)

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return
	}
	
	page.User_id = user.Id
	page, err = pageController.pageService.Store(page)

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))	
		return
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Página criada com sucesso", page))
}

func (pageController *PageController) Delete(context *gin.Context){
	slug := context.Param("slug")

	if slug == ""{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "Slug não foi encontrado!", nil))
	}
	
	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	err = pageController.pageService.Delete(slug, user.Id)

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Página deletada com sucesso!", nil))
}

func (pageController *PageController) UpdateTitle (context *gin.Context){
	slug := context.Param("slug")

	if slug == ""{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "Slug não foi encontrado!", nil))
	}

	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	var titleRequest requests.TitleRequest
	err = context.ShouldBindJSON(&titleRequest)

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	title, slug, err := pageController.pageService.UpdateTitle(titleRequest.Title, slug, user.Id)
	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Título atualizado com sucesso!", gin.H{"title": title, "slug": slug}))
}

func (pageController *PageController) UpdateEmoji (context *gin.Context){
	slug := context.Param("slug")

	if slug == ""{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "Slug não foi encontrado!", nil))
	}

	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	var emojiRequest requests.EmojiRequest
	err = context.ShouldBindJSON(&emojiRequest)

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	emoji, err := pageController.pageService.UpdateEmoji(emojiRequest.Emoji, slug, user.Id)
	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Emoji atualizado com sucesso!", gin.H{"emoji": emoji, "slug": slug}))
}

func (pageController *PageController) UpdateContent(context *gin.Context){
	slug := context.Param("slug")

	var contentRequest requests.ContentRequest

	if slug == ""{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "Slug não foi encontrado!", nil))
	}

	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	err = context.ShouldBindJSON(&contentRequest)

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return		
	}
	
	err = pageController.pageService.UpdateContent(contentRequest.Content, slug, user.Id)
	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "Conteúdo atualizado com sucesso!", true))
}