package router

import (
	"github.com/diogenesOliver/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpenningHandler)
		v1.POST("/opening", handler.CreateOpenningHandler)
		v1.DELETE("/opening", handler.DeleteOpenningHandler)
		v1.PUT("/opening", handler.ShowOpenningHandler)
		v1.GET("/openings", handler.UpdateOpenningHandler)
	}
}
