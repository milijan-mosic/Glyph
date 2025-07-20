package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	status := http.StatusNoContent

	response := HelloResponse{
		Message: "Article is deleted...",
		Status:  status,
	}

	c.JSON(status, response)
}
