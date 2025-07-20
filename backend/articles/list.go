package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	status := http.StatusOK

	response := HelloResponse{
		Message: "List of articles :)",
		Status:  status,
	}

	c.JSON(status, response)
}
