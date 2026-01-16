package article

import (
	"encoding/json"
	"log"
	"net/http"

	"glyph/database"
	"glyph/models"

	"gorm.io/gorm"
)

type UpdateArticleRequest struct {
	ArticleID   string `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	Content     string `json:"content"`
}

func Update(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateArticleRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			database.Error(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}

		if req.ArticleID == "" ||
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
			Model(&models.Article{}).
			Where("id = ?", req.ArticleID).
			Updates(updates)

		if result.Error != nil {
			log.Printf("Error while updating article: %v", result.Error)
			database.Error(w, http.StatusInternalServerError, result.Error.Error())
			return
		}

		if result.RowsAffected == 0 {
			database.Error(w, http.StatusNotFound, "Article not found")
			return
		}

		database.JSON(w, http.StatusOK, "message", "Article updated!")
	}
}
