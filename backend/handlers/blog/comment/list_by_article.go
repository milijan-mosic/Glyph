package comment

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func ListByArticleID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentID := chi.URLParam(r, "commentId")
		if commentID == "" {
			database.Error(w, http.StatusBadRequest, "Comment ID is required")
			return
		}

		var comments []models.Comment

		result := db.
			Order("created_at DESC").
			Find(&comments)

		err := result.Error

		if err != nil {
			log.Printf("Error while listing comments: %v", err)

			database.JSON(w, http.StatusInternalServerError, "message", err)
			return
		}

		if len(comments) == 0 {
			comments = make([]models.Comment, 0)
		}

		database.JSON(w, http.StatusOK, "comments", comments)
	}
}
