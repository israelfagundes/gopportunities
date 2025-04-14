package router

import "github.com/gin-gonic/gin"

func Init() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initRoutes(router)

	// Run the server
	router.Run()
}
