package main

import (
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

	comment := router.Group(urlPrefix + "/comment")
	{
		comment.POST("/create", comments.CreateComment)
		comment.GET("/list/:articleId", comments.GetApprovedComments)
		comment.GET("/pending", comments.GetPendingComments)
		comment.PUT("/approve/:commentId", comments.ApproveComment)
		comment.DELETE("/delete/:commentId", comments.DeleteComment)
	}

	log.Printf("Server starting on: http://localhost%s\n", address)
	router.Run(address)
}
