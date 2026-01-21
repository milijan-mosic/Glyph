package comment

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"glyph/database"
	"glyph/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateCommentRequest struct {
	ArticleID  string `json:"article_id" binding:"required"`
	AuthorName string `json:"author_name" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

func Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCommentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			database.Error(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}

		if req.ArticleID == "" || req.Content == "" || req.AuthorName == "" {
			database.Error(w, http.StatusBadRequest, "Arguments `article_id`, `content` and `author_name` are required!")
			return
		}

		comment := models.Comment{
			ID:         uuid.NewString(),
			ArticleID:  req.ArticleID,
			AuthorName: req.AuthorName,
			Content:    req.Content,
			Approved:   false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		result := db.Create(&comment)

		err := result.Error

		if err != nil {
			log.Printf("Error while creating comment: %v", err)
			database.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		database.JSON(w, http.StatusCreated, "commentId", comment.ID)
	}
}
