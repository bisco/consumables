package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConsumableItem struct {
	Name            string `json:"name"`
	Count           uint32 `json:"count"`
	Id              uint32 `json:"id"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
	Category        string `json:"category"`
	CategoryCode    string `json:"category_code"`
	SubCategory     string `json:"subcategory"`
	SubCategoryCode string `json:"subcategory_code"`
}

type CategoryCode struct {
	CategoryCode    string `json:"category_code"`
	SubCategoryCode string `json:"subcategory_code"`
}

// string -> hex
var categorycodeCache = map[string]string{}

func genCategoryCode(s string) string {
	v, ok := categorycodeCache[s]
	if ok {
		return v
	} else {
		categorycodeCache[s] = fmt.Sprintf("%x", s)
		return categorycodeCache[s]
	}
}

// PATH: /
func genIndexView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cs := dbGetAll()
		// categories["category"]["subcategory"] = {categorycode, subcategorycode}
		var categories = map[string]map[string]CategoryCode{}
		var allItems []ConsumableItem
		for _, c := range cs {
			allItems = append(allItems,
				ConsumableItem{
					Name:            c.Name,
					Id:              c.Id,
					Count:           c.Count,
					UpdatedAt:       c.UpdatedAt.UnixNano() / 1000000,
					CreatedAt:       c.CreatedAt.UnixNano() / 1000000,
					Category:        c.Category,
					SubCategory:     c.SubCategory,
					CategoryCode:    genCategoryCode(c.Category),
					SubCategoryCode: genCategoryCode(c.SubCategory),
				})
			_, ok := categories[c.Category]
			if !ok {
				categories[c.Category] = map[string]CategoryCode{}
			}
			categories[c.Category][c.SubCategory] = CategoryCode{CategoryCode: genCategoryCode(c.Category),
				SubCategoryCode: genCategoryCode(c.SubCategory)}
		}
		ctx.HTML(http.StatusOK, "index.html",
			gin.H{"allItems": allItems, "categoryCodeMap": categories})
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
		idstr := ctx.Param("id")
		id, err := strconv.ParseUint(idstr, 10, 32)
		if err != nil {
			log.Fatalf("invalid input", err)
		}
		c := dbGetById(uint32(id))
		ctx.HTML(http.StatusOK, "modify.html",
			gin.H{
				"item": c,
			})
	}
}
