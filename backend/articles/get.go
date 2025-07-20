package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticleRoute(c *gin.Context) {
	status := http.StatusOK

	response := HelloResponse{
		Message: "Single article!",
		Status:  status,
	}

	c.JSON(status, response)
}
