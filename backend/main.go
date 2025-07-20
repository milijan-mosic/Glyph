package main

import (
	"heartbit/articles"
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

		article.GET("/list", articles.ListArticle)
		article.GET("/get", articles.GetArticle)
		article.POST("/create", articles.CreateArticle)
		article.PUT("/update", articles.UpdateArticle)
		article.DELETE("/delete", articles.DeleteArticle)
	}

	log.Printf("Server starting on: http://localhost%s\n", address)
	router.Run(address)
}
