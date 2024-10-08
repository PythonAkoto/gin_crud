package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pythonakoto/gin_crud/controllers"
	"github.com/pythonakoto/gin_crud/repositories"
	"github.com/pythonakoto/gin_crud/services"
)

func InitialiseRoutes(router *gin.Engine) {
	// initaslise repositories and services
	itemRepo := repositories.NewItemRepository()
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	// CRUD routes
	router.GET("/items", itemController.GetItems)
	router.GET("/items/:id", itemController.GetItem)
	router.POST("/items", itemController.CreateItem)
	router.PUT("/items/:id", itemController.UpdateItem)
	router.DELETE("/items/:id", itemController.DeleteItem)

	// HTML route to render items list
	router.GET("/items/html", itemController.RenderItemsHTML)
}
