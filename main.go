package main

import (
	"context"
	"net/http"

	"github.com/conneroisu/go-htmx-templ-tailwind-template/internal/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	var list []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen"}
	component := pages.Home(list)
	runner := gin.Default()
	runner.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	runner.GET("/", func(c *gin.Context) {
		component.Render(context.Background(), c.Writer)
	})
	runner.Run()
}
