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
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	Content     string `json:"content"`
}

func Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCommentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			database.Error(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}

		if req.Title == "" || req.Content == "" {
			database.Error(w, http.StatusBadRequest, "Arguments `title` and `content` are required!")
			return
		}

		comment := models.Comment{
			ID:          uuid.NewString(),
			Title:       req.Title,
			Description: req.Description,
			Content:     req.Content,
			Published:   req.Published,
			Author:      "Milijan Mosic",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		result := db.Create(&comment)

		err := result.Error

		if err != nil {
			log.Printf("Error while creating comment: %v", err)
			database.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		database.JSON(
			w,
			http.StatusCreated,
			"commentId",
			comment.ID,
		)
	}
}
