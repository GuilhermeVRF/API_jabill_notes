package controllers

import (
	"jabill-notes/auth"
	"jabill-notes/models"
	"jabill-notes/services"
	"jabill-notes/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
    return UserController{
        userService: userService,
    }
}

func (userController *UserController) Index (context *gin.Context){
	users, err := userController.userService.Index();
	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), []interface{}{}))	
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "User retrieved!", users))
}

func (userController *UserController) Show (context *gin.Context){
	authorizationHeader := context.GetHeader("Authorization")

	user, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return 
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "User retrieved!", user))
}

func (userController *UserController) Store(context *gin.Context) {
	err := context.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return
	}

	name := context.PostForm("name")
	email := context.PostForm("email")
	password := context.PostForm("password")
	repeatPassword := context.PostForm("repeat_password")

	// Validação das senhas
	if password != repeatPassword {
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", "As senhas não coincidem!", nil))
		return
	}

	file, header, err := context.Request.FormFile("profile_image")

	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
		return
	}
	defer file.Close()
	
	fileUploader := utils.NewFilesUploader("./uploads/users/")
	filePath, err := fileUploader.SaveFile(file, header)

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return
	}

	user := models.User{
		Name:         name,
		Email:        email,
		Password:     password,
		Profile_image: filePath,
	}

	user, err = userController.userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return
	}
	userToken, err := user.GenerateToken()

	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return	
	}

	context.JSON(http.StatusCreated, utils.NewApiResponse("success", "Usuário cadastrado com sucesso!", gin.H{"user_token" : userToken }))
}

func (userController *UserController) Update(context *gin.Context) {
	var user models.User
	authorizationHeader := context.GetHeader("Authorization")

	parsedUser, err := auth.ParseToken(authorizationHeader)

	if err != nil{
		context.JSON(http.StatusUnauthorized, utils.NewApiResponse("error", "Usuário não autenticado!", nil))
		return
	}

	err = context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
	}

	user.Id = parsedUser.Id
	user, err = userController.userService.UpdateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewApiResponse("error", err.Error(), nil))
		return
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "User updated!", user))
}

func (userController *UserController) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))	
	}

	_, err = userController.userService.Show(id)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.NewApiResponse("error", "User with Id "+ strconv.Itoa(id) + " not found!", nil))
		return
	}

	_, err = userController.userService.DeleteUser(id)
	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewApiResponse("error", err.Error(), nil))
	}

	context.JSON(http.StatusOK, utils.NewApiResponse("success", "User deleted!", true))
}