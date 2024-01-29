package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST lanchonete-de-bairro",
	})
}
