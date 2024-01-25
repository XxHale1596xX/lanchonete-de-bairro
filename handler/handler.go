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

func ShowLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET lanchonete-de-bairro",
	})
}

func DeleteLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE lanchonete-de-bairro",
	})
}

func UpdateLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE lanchonete-de-bairro",
	})
}

func ListLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT lanchonete-de-bairro",
	})
}

func EditLanchoneteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PATH lanchonete-de-bairro",
	})
}
