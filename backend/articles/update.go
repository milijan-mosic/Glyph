package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateArticle(c *gin.Context) {
	status := http.StatusOK

	response := HelloResponse{
		Message: "Updated article.",
		Status:  status,
	}

	c.JSON(status, response)
}
