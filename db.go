package main

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Consumable struct {
	Id        uint32    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Count     uint32    `json:"count"`
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

func dbInsert(name string, count uint32) {
	db := dbOpen()
	db.Create(&Consumable{Name: name, Count: count})
}

func dbUpdate(id uint32, name string, count uint32) {
	db := dbOpen()

	var c Consumable
	db.First(&c, id)
	c.Name = name
	c.Count = count
	db.Save(&c)
}

func dbDelete(id uint32) {
	db := dbOpen()

	var c Consumable
	db.Delete(&c, id)
}

func dbGetAll() []Consumable {
	db := dbOpen()

	var cs []Consumable
	db.Order("updated_at asc").Find(&cs)
	return cs
}

//-----------------------------------
// test purpose only
//-----------------------------------
func dbDataGen() {
	for i := 0; i < 10; i++ {
		dbInsert("test"+strconv.Itoa(i), uint32(i))
	}
}
