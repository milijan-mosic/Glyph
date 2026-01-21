package comment

import (
	"encoding/json"
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"gorm.io/gorm"
)

type UpdateCommentRequest struct {
	CommentID   string `json:"comment_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	Content     string `json:"content"`
}

func Approve(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateCommentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			database.Error(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}

		if req.CommentID == "" ||
			req.Title == "" ||
			req.Description == "" ||
			req.Content == "" {
			database.Error(w, http.StatusBadRequest, "all fields are required")
			return
		}

		updates := map[string]any{
			"title":       req.Title,
			"description": req.Description,
			"content":     req.Content,
			"published":   req.Published,
		}

		result := db.
			Model(&models.Comment{}).
			Where("id = ?", req.CommentID).
			Updates(updates)

		err := result.Error

		if err != nil {
			log.Printf("Error while updating comment: %v", err)
			database.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		if result.RowsAffected == 0 {
			database.Error(w, http.StatusNotFound, "Comment not found")
			return
		}

		database.JSON(w, http.StatusOK, "message", "Comment updated!")
	}
}
