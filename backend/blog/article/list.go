package articles

import (
	"context"
	"glyph/database_interfaces"
	db "glyph/db"
	"glyph/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ListArticleRoute(c *gin.Context) {
	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	operations := database_interfaces.New(conn)

	articles, err := operations.ListArticles(ctx)
	if err != nil {
		log.Printf("Error while listing articles: %s", err)

		status := http.StatusInternalServerError
		response := db.Response{
			Message: "Articles listing failed!",
			Status:  status,
		}
		c.JSON(status, response)
	}

	if len(articles) == 0 {
		articles = make([]database_interfaces.Article, 0)
	}
	status := http.StatusOK
	c.JSON(status, gin.H{"articles": articles})
}
