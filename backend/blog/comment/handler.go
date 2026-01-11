package comment

import "glyph/database_interfaces"

type CommentHandler struct {
	queries *database_interfaces.Queries
}

func NewCommentHandler(q *database_interfaces.Queries) *CommentHandler {
	return &CommentHandler{queries: q}
}
