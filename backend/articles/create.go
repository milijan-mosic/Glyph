package articles

import (
	"context"
	"heartbit/database_interfaces"
	"heartbit/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateArticleRoute(c *gin.Context) {
	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	operations := database_interfaces.New(conn)

	article, err := operations.CreateArticle(ctx, database_interfaces.CreateArticleParams{
		ID:        uuid.NewString(),
		Title:     "Test title :)",
		Author:    "Milijan Mosic",
		Content:   "# Hello!\n",
		Published: true,
	})
	if err != nil {
		log.Printf("Error while creating article: %s", err)

		status := http.StatusInternalServerError
		response := HelloResponse{
			Message: "Article creation failed!",
			Status:  status,
		}
		c.JSON(status, response)
	}

	log.Printf("Created new article %v", article)

	status := http.StatusCreated
	response := HelloResponse{
		Message: "Article created!",
		Status:  status,
	}
	c.JSON(status, response)
}
