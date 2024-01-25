package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//show lanchonete de bairro
		v1.GET("/lanchonete-de-bairro", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Lanchonete-de-bairro",
			})
		})
		v1.POST("/lanchonete-de-bairro", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Lanchonete-de-bairro",
			})
		})
		v1.DELETE("/lanchonete-de-bairro", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Lanchonete-de-bairro",
			})
		})
		v1.PUT("/lanchonete-de-bairro", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Lanchonete-de-bairro",
			})
		})
		v1.PATCH("/lanchonete-de-bairro", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PATCH Lanchonete-de-bairro",
			})
		})

	}
}
