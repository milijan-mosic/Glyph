package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")

	if err := h.queries.DeleteComment(c, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
