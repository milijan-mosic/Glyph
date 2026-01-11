package articles

import (
	"context"
	"glyph/database_interfaces"
	db "glyph/db"
	"glyph/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateArticleArguments struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	Content     string `json:"content" binding:"required"`
}

func CreateArticleRoute(c *gin.Context) {
	var arguments CreateArticleArguments
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

	articleId := uuid.NewString()

	article, err := operations.CreateArticle(ctx, database_interfaces.CreateArticleParams{
		ID:          articleId,
		Title:       arguments.Title,
		Author:      "Milijan Mosic",
		Description: arguments.Description,
		Content:     arguments.Content,
		Published:   arguments.Published,
		ModifiedAt:  timestamp,
		CreatedAt:   timestamp,
	})
	if err != nil {
		log.Printf("Error while creating article: %s", err)

		status := http.StatusInternalServerError
		response := db.Response{
			Message: "Article creation failed!",
			Status:  status,
		}
		c.JSON(status, response)
	}

	log.Printf("Created new article %v", article)

	status := http.StatusCreated
	c.JSON(status, gin.H{"ArticleId": articleId})
}
