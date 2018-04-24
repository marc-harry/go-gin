package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductCreate create a product
func ProductCreate(c *gin.Context) {
	var product Product
	if err := c.ShouldBind(&product); err == nil {
		db, dbErr := GetDb()
		defer db.Close()
		if dbErr == nil {
			db.Create(&product)
			c.JSON(http.StatusOK, product)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": dbErr.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
}

// ProductGet all products
func ProductGet(c *gin.Context) {
	if db, err := GetDb(); err == nil {
		defer db.Close()
		var products []Product
		db.Find(&products)
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
	}
}

// ProductDelete a product by id
func ProductDelete(c *gin.Context) {
	id, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": idErr.Error()})
	}
	if db, err := GetDb(); err == nil {
		defer db.Close()
		var product Product
		db.First(&product, id)
		if product.ID != 0 {
			db.Delete(&product)
			c.JSON(http.StatusNoContent, nil)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": fmt.Sprintf("No product found with Id: %d", id)})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
	}
}
