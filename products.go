package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

// ProductControllerType interface
type ProductControllerType struct{}

// ProductController instance
var ProductController = ProductControllerType{}

// Create a new product
func (ctrl *ProductControllerType) Create(c *gin.Context) {
	var product Product
	if err := ctrl.shouldBind(c, &product); err == nil {
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

// Get all products
func (ctrl *ProductControllerType) Get(c *gin.Context) {
	log.Info("Get Products Called")
	if db, err := GetDb(); err == nil {
		defer db.Close()
		var products []Product
		db.Find(&products)
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
	}
}

// Delete a product by id
func (ctrl *ProductControllerType) Delete(c *gin.Context) {
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

func (ctrl *ProductControllerType) shouldBind(c *gin.Context, obj interface{}) error {
	return c.ShouldBindWith(obj, binding.Default(c.Request.Method, c.ContentType()))
}
