package articles

import (
	"context"
	"heartbit/database_interfaces"
	"heartbit/utils"
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

	queries := database_interfaces.New(conn)

	articles, err := queries.ListArticles(ctx)
	if err != nil {
		log.Fatalf("Error while listing articles: %s", err)

		status := http.StatusInternalServerError
		response := HelloResponse{
			Message: "Articles listing failed!",
			Status:  status,
		}
		c.JSON(status, response)
	}

	status := http.StatusOK
	c.JSON(status, articles)
}
