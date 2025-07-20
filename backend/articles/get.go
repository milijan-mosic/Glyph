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

func GetArticleRoute(c *gin.Context) {
	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	operations := database_interfaces.New(conn)
	articleId := c.Param("articleId")

	article, err := operations.GetArticle(ctx, articleId)
	if err != nil {
		log.Printf("Error while getting article: %s", err)
	}

	status := http.StatusOK
	c.JSON(status, gin.H{"article": article})
}
