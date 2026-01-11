package comments

import (
	"context"
	"glyph/database_interfaces"
	"glyph/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createCommentRequest struct {
	ArticleId     string  `json:"article_id" binding:"required"`
	AuthorName string `json:"author_name" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

func CreateComment(c *gin.Context) {
	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	operations := database_interfaces.New(conn)

	comment, err := operations.CreateComment(c, database_interfaces.CreateCommentParams{
		ArticleID:     req.ArticleId,
		AuthorName: req.AuthorName,
		Content:    req.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
