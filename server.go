package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TargetId struct {
	Id uint32 `json:"id"`
}

type CreateReq struct {
	Name        string `json:"name"`
	Count       uint32 `json:"count"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

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
	r.POST(V1_API_PREFIX+"/consumable-items/actions/plus-one/invoke", countPlusOne())
	r.POST(V1_API_PREFIX+"/consumable-items/actions/minus-one/invoke", countMinusOne())
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
		var reqbody CreateReq
		if err := ctx.ShouldBindJSON(&reqbody); err != nil {
			log.Fatalf("bind fail: %v", err)
		}
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			log.Fatalf("fail parseUint: %v", err)
		}
		dbUpdate(uint32(id), reqbody.Name, reqbody.Count, reqbody.Category, reqbody.SubCategory)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})

	}
}

func createItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqbody CreateReq
		if err := ctx.ShouldBindJSON(&reqbody); err != nil {
			log.Fatalf("bind fail: %v", err)
		}
		dbInsert(reqbody.Name, reqbody.Count, reqbody.Category, reqbody.SubCategory)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func deleteItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			log.Fatalf("fail parseUint: %v", err)
		}
		dbDelete(uint32(id))
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func countPlusOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body TargetId
		if err := ctx.ShouldBindJSON(&body); err != nil {
			log.Fatalf("bind fail: %v", err)
		}
		c := dbGetById(body.Id)
		c.Count += 1
		dbModifyCount(c.Id, c.Name, c.Count)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func countMinusOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body TargetId
		if err := ctx.ShouldBindJSON(&body); err != nil {
			log.Fatalf("bind fail: %v", err)
		}
		c := dbGetById(body.Id)
		if c.Count == 0 {
			ctx.JSON(http.StatusOK, gin.H{"status": "fail"})
		}
		c.Count -= 1
		dbModifyCount(c.Id, c.Name, c.Count)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
