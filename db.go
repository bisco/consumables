package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Consumable struct {
	Id          uint32    `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Count       uint32    `json:"count"`
	Category    string    `json:"category"`
	SubCategory string    `json:"subcategory"`
}

const DB_PATH = "./consumables.db"

var dbConn *gorm.DB = dbConnInit()

func dbConnInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func dbOpen() *gorm.DB {
	return dbConn
}

func dbInit() {
	fmt.Println("dbInit")
	db := dbOpen()
	db.AutoMigrate(&Consumable{})
}

func dbInsert(name string, count uint32, category, subcategory string) {
	db := dbOpen()
	result := db.Create(&Consumable{Name: name, Count: count, Category: category, SubCategory: subcategory})
	fmt.Println(result)
}

func dbUpdate(id uint32, name string, count uint32, category, subcategory string) {
	db := dbOpen()
	fmt.Println("dbUpdate:", id, name, count, category, subcategory)
	var c Consumable
	db.First(&c, id)
	c.Name = name
	c.Count = count
	c.Category = category
	c.SubCategory = subcategory
	db.Save(&c)
}

func dbModifyCount(id uint32, name string, count uint32) {
	db := dbOpen()
	fmt.Println("dbModifyCount:", id, name, count)
	var c Consumable
	db.First(&c, id)
	c.Name = name
	c.Count = count
	db.Save(&c)
}

func dbDelete(id uint32) {
	db := dbOpen()

	var c Consumable
	result := db.Delete(&c, id)
	fmt.Println(c, result)
}

func dbGetAll() []Consumable {
	db := dbOpen()

	var cs []Consumable
	db.Order("count asc").Find(&cs)
	return cs
}

func dbGetById(id uint32) Consumable {
	db := dbOpen()

	var c Consumable
	db.First(&c, id)
	return c
}
