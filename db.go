package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// Product item
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// GetDb Gets DB instance to operate with DB
func GetDb() (*gorm.DB, error) {
	db, err := gorm.Open("mssql", "sqlserver://GoUser:password@localhost:1433?database=go-test")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Product{})
	return db, err
}
