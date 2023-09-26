package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router
	router := gin.Default()
	
	//Initialize routes
	initializeRoutes(router)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run() 
}