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
	CommentID string `json:"comment_id"`
	Approved  bool   `json:"approved"`
}

func Approve(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateCommentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			database.Error(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}

		if req.CommentID != "" {
			database.Error(w, http.StatusBadRequest, "All fields are required")
			return
		}

		updatedComment := map[string]any{
			"approved": req.Approved,
		}

		result := db.
			Model(&models.Comment{}).
			Where("id = ?", req.CommentID).
			Updates(updatedComment)

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

		database.JSON(w, http.StatusNoContent, "", "")
	}
}
