-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at;

-- name: GetArticle :one
SELECT * FROM articles
WHERE id = $1
LIMIT 1;

-- name: CreateArticle :one
INSERT INTO articles (
  id, title, description, content, author, published, created_at, modified_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateArticle :exec
UPDATE articles
SET title = $1, description = $2, content = $3, published = $4, modified_at = $5
WHERE id = $6
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;
