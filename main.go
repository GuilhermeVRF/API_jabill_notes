package main

import (
	"jabill-notes/controllers"
	"jabill-notes/database"
	"jabill-notes/repository"
	"jabill-notes/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},     
	}))

	dataBaseConnection, databaseError  := database.Connect()
	if(databaseError != nil){
		panic(databaseError)
	}

	UserRepository := repository.NewUserRepository(dataBaseConnection)
	UserService := services.NewUserService(UserRepository)
	UserController := controllers.NewUserController(UserService)

	server.Static("./uploads", "/uploads")

	// User routes
	server.GET("/users", UserController.Index)
	server.GET("/user", UserController.Show)
	server.POST("/user", UserController.Store)
	server.PUT("/user/:id", UserController.Update)
	server.DELETE("/user/:id", UserController.Delete)

	LoginRepository := repository.NewLoginRepository(dataBaseConnection)
	LoginService := services.NewLoginService(LoginRepository)
	LoginController := controllers.NewLoginController(LoginService)

	// Login Routes
	server.POST("/login", LoginController.Verify)

	MediaRepository := repository.NewMediaRepository(dataBaseConnection)
	MediaService := services.NewMediaService(MediaRepository)
	MediaController := controllers.NewMediaController(MediaService)

	server.GET("/user/profile", MediaController.Show)

	PageRepository := repository.NewPageRepository(dataBaseConnection)
	PageService := services.NewPageService(PageRepository)
	PageController := controllers.NewPageController(PageService)

	// Page routes
	server.GET("/page/:slug", PageController.Show)
	server.GET("/pages", PageController.Index)
	server.POST("/page", PageController.Store)
	server.PUT("/page/content/:slug", PageController.UpdateContent)
	server.PUT("/page/title/:slug", PageController.UpdateTitle)
	server.PUT("/page/emoji/:slug", PageController.UpdateEmoji)
	server.DELETE("/page/:slug", PageController.Delete)

	server.Run(":5000")
}
