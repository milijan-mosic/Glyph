package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CommentHandler) GetApprovedComments(c *gin.Context) {
	articleId := c.Param("articleId")

	comments, err := h.queries.GetApprovedCommentsByPost(c, articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
