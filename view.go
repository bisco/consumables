package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PATH: /
func genIndexView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allItems := dbGetAll()
		ctx.HTML(http.StatusOK, "index.html",
			gin.H{"allItems": allItems})
	}
}

// PATH: /add
func genAddView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "add.html",
			gin.H{})
	}
}

// PATH: /modify/:id
func genModifyView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//n := ctx.Param("id")
	}
}
