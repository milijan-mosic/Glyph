package article

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleId")
		if articleID == "" {
			database.Error(w, http.StatusBadRequest, "Article ID is required")
			return
		}

		result := db.
			Where("id = ?", articleID).
			Delete(&models.Article{})

		if result.Error != nil {
			log.Printf("Error while deleting article: %v", result.Error)
			database.Error(w, http.StatusInternalServerError, result.Error.Error())
			return
		}

		if result.RowsAffected == 0 {
			database.Error(w, http.StatusNotFound, "Article not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
