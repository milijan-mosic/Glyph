package main

import (
	articles "glyph/blog/article"
	comments "glyph/blog/comment"
	"glyph/db"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	address   = ":5000"
	urlPrefix = "/api/1.0"
)

func main() {
	router := gin.Default()
	
	db.InitializeDatabase()

	article := router.Group(urlPrefix + "/article")
	{
		article.GET("/list", articles.ListArticleRoute)
		article.GET("/get/:articleId", articles.GetArticleRoute)
		article.POST("/create", articles.CreateArticleRoute)
		article.PUT("/update", articles.UpdateArticleRoute)
		article.DELETE("/delete/:articleId", articles.DeleteArticleRoute)
	}

	comment := router.Group(urlPrefix + "/comment")
	{
		comment.POST("/comments", comments.CreateComment)
		comment.GET("/posts/:postId/comments", comments.GetApprovedComments)
		comment.GET("/admin/comments/pending", comments.GetPendingComments)
		comment.PUT("/admin/comments/:id/approve", comments.ApproveComment)
		comment.DELETE("/admin/comments/:id", comments.DeleteComment)
	}

	log.Printf("Server starting on: http://localhost%s\n", address)
	router.Run(address)
}
