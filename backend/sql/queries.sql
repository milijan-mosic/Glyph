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

----------------------------------------------------------------

-- name: CreateComment :one
INSERT INTO comments (id, article_id, author_name, content)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetApprovedCommentsByPost :many
SELECT *
FROM comments
WHERE article_id = $1 AND approved = TRUE
ORDER BY created_at ASC;

-- name: GetPendingComments :many
SELECT *
FROM comments
WHERE approved = FALSE
ORDER BY created_at ASC;

-- name: ApproveComment :exec
UPDATE comments
SET approved = TRUE
WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;
