package article

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func GetByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articleId := chi.URLParam(r, "articleId")
		article := models.Article{}

		result := db.
			Where("id = ?", articleId).
			Find(&article)

		err := result.Error

		if err != nil {
			log.Printf("Error while listing article: %v", err)

			database.JSON(w, http.StatusInternalServerError, "message", err)
			return
		}

		database.JSON(w, http.StatusOK, "article", article)
	}
}
