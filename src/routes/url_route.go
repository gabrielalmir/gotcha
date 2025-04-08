package routes

import (
	"gotcha/src/controllers"
	"gotcha/src/db"
	"gotcha/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	database := db.NewDatabase("mongodb://localhost:27017")
	if err := database.Connect(); err != nil {
		panic(err)
	}

	// Initialize services
	urlService, err := services.NewURLService(database)
	if err != nil {
		panic(err)
	}

	// Initialize controllers
	urlController := controllers.NewURLController(urlService)

	// URL routes
	router.POST("/shorten-url", urlController.CreateShortURL)
	router.GET("/:short", urlController.RedirectToOriginal)
}
