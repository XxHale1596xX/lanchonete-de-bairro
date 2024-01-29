package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PATH lanchonete-de-bairro",
	})
}
