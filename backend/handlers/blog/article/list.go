package article

import (
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"gorm.io/gorm"
)

func List(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var articles []models.Article

		result := db.
			Order("created_at DESC").
			Find(&articles)

		err := result.Error

		if err != nil {
			log.Printf("Error while listing articles: %v", err)

			database.JSON(w, http.StatusInternalServerError, "message", err)
			return
		}

		if len(articles) == 0 {
			articles = make([]models.Article, 0)
		}

		database.JSON(w, http.StatusOK, "articles", articles)
	}
}
