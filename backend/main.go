package main

import (
	"glyph/articles"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	address   = ":5000"
	urlPrefix = "/api/1.0"
)

func main() {
	router := gin.Default()

	{
		articles.InitializeDatabase()
		article := router.Group(urlPrefix + "/article")

		article.GET("/list", articles.ListArticleRoute)
		article.GET("/get/:articleId", articles.GetArticleRoute)
		article.POST("/create", articles.CreateArticleRoute)
		article.PUT("/update", articles.UpdateArticleRoute)
		article.DELETE("/delete/:articleId", articles.DeleteArticleRoute)
	}

	log.Printf("Server starting on: http://localhost%s\n", address)
	router.Run(address)
}
