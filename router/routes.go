package router

import (
	"github.com/XxHale1596xX/lanchonete-de-bairro/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//show lanchonete de bairro
		v1.GET("/lanchonete-de-bairro", handler.ShowLanchoneteHandler)
		v1.POST("/lanchonete-de-bairro", handler.CreateLanchoneteHandler)
		v1.DELETE("/lanchonete-de-bairro", handler.DeleteLanchoneteHandler)
		v1.PUT("/lanchonete-de-bairro", handler.ListLanchoneteHandler)
		v1.PATCH("/lanchonete-de-bairro", handler.UpdateLanchoneteHandler)

	}
}
