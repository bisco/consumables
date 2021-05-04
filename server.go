package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	// gen html
	r.GET("/", genIndexView())
	r.GET("/add", genAddView())
	r.GET("/modify/:id", genModifyView())

	// API Ver.1
	const V1_API_PREFIX = "/api/v1"
	r.GET(V1_API_PREFIX+"/consumable-items", getAllItems())
	r.GET(V1_API_PREFIX+"/consumable-items/:id", getItemById())
	r.PATCH(V1_API_PREFIX+"/consumable-items/:id", modifyItem())
	r.POST(V1_API_PREFIX+"/consumable-items", createItem())
	r.DELETE(V1_API_PREFIX+"/consumable-items/:id", deleteItem())

	return r
}

func getAllItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cs := dbGetAll()
		ctx.JSON(http.StatusOK, cs)
	}
}

func getItemById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func modifyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func createItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func deleteItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
