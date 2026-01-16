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

func ApproveComment(c *gin.Context) {
	commentId := c.Param("commentId")

	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	operations := database_interfaces.New(conn)

	if err := operations.ApproveComment(c, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
