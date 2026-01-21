package comment

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"gorm.io/gorm"
)

func ListPending(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var comments []models.Comment

		result := db.
			Where("approved = ?", false).
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
