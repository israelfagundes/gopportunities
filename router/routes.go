package router

import (
	"github.com/gin-gonic/gin"
	"github.com/israelfagundes/gopportunities/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
