package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumableItem struct {
	Name      string `json:"name"`
	Count     uint32 `json:"count"`
	Id        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// PATH: /
func genIndexView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cs := dbGetAll()
		var allItems []ConsumableItem
		for _, c := range cs {
			allItems = append(allItems,
				ConsumableItem{
					Name:      c.Name,
					Id:        c.Id,
					Count:     c.Count,
					UpdatedAt: c.UpdatedAt.UnixNano() / 1000000,
					CreatedAt: c.CreatedAt.UnixNano() / 1000000,
				})
		}
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
