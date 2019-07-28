package main

import (
	"github.com/gin-gonic/gin"
	"server/blog"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	group := r.Group("/api/v1/")

	group.GET("/blog/search", blog.Search)
	group.GET("/blog/favor", blog.Favor)
	group.GET("/blog/bookmarks", blog.Bookmarks)

	r.Run() // listen and serve on 0.0.0.0:8080
}
