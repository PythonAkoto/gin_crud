package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pythonakoto/gin_crud/routes"
)

func main() {
	router := gin.Default()

	// load HTML templates from the templates folder
	router.LoadHTMLGlob("templates/*") // load all from templates

	// optional: serve any static files you may have --> css, js, etc
	router.Static("/static", "./static")

	//initialise the routes
	routes.InitialiseRoutes(router)

	// start the server
	router.Run(":8080")
}
