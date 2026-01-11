package articles

import (
	"context"
	"glyph/database_interfaces"
	"glyph/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateArticleArguments struct {
	ArticleId   string `json:"article_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Published   bool   `json:"published" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func UpdateArticleRoute(c *gin.Context) {
	var arguments UpdateArticleArguments
	if err := c.ShouldBindJSON(&arguments); err != nil {
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

	var timestamp pgtype.Timestamp
	timestamp.Time = time.Now()
	timestamp.Valid = true

	err = operations.UpdateArticle(ctx, database_interfaces.UpdateArticleParams{
		Title:       arguments.Title,
		Description: arguments.Description,
		Content:     arguments.Content,
		Published:   arguments.Published,
		ModifiedAt:  timestamp,
		ID:          arguments.ArticleId,
	})
	if err != nil {
		log.Printf("Error while updating article: %s", err)

		status := http.StatusInternalServerError
		response := HelloResponse{
			Message: "Update article failed!",
			Status:  status,
		}
		c.JSON(status, response)
	}

	log.Print("Updated article...")

	status := http.StatusOK
	response := HelloResponse{
		Message: "Article updated!",
		Status:  status,
	}
	c.JSON(status, response)
}
