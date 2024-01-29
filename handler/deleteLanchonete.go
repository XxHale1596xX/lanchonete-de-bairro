package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE lanchonete-de-bairro",
	})
}
