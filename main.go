package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/*")

	api := router.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "from api success",
			})
		})

		productRoute := api.Group("/product")
		{
			productRoute.POST("/create", ProductCreate)
			productRoute.GET("/all", ProductGet)
			productRoute.DELETE("/:id", ProductDelete)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/")
	})

	router.Run()
}
