package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	var library []Book

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		book, err := createBook()
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
			return
		}
		library = append(library, book)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
