package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE lanchonete-de-bairro",
	})
}
