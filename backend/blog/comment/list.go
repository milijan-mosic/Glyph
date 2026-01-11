package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CommentHandler) GetPendingComments(c *gin.Context) {
	comments, err := h.queries.GetPendingComments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}