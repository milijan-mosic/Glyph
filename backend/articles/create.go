package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	status := http.StatusCreated

	response := HelloResponse{
		Message: "Article created!",
		Status:  status,
	}

	c.JSON(status, response)
}
