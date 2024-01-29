package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT lanchonete-de-bairro",
	})
}
