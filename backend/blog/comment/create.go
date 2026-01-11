package comment

import (
	"glyph/database_interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCommentRequest struct {
	ArticleId     string  `json:"article_id" binding:"required"`
	AuthorName string `json:"author_name" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := h.queries.CreateComment(c, database_interfaces.CreateCommentParams{
		ArticleID:     req.ArticleId,
		AuthorName: req.AuthorName,
		Content:    req.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
