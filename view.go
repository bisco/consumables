package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PATH: /
func genIndexView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allItems := dbGetAll()
		fmt.Println(allItems)
		ctx.HTML(http.StatusOK, "index.html",
			gin.H{"allItems": allItems})
	}
}

// PATH: /add
func genAddView() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// PATH: /modify/:id
func genModifyView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//n := ctx.Param("id")
	}
}
