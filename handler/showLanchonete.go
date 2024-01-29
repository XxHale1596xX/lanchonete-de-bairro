package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET lanchonete-de-bairro",
	})
}
