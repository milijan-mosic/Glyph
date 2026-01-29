package comment

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Approve(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentID := chi.URLParam(r, "commentId")
		if commentID == "" {
			database.Error(w, http.StatusBadRequest, "Comment ID is required")
			return
		}

		updatedComment := map[string]any{
			"approved": true,
		}

		result := db.
			Model(&models.Comment{}).
			Where("id = ?", commentID).
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
