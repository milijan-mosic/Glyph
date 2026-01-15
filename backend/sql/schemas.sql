CREATE TABLE articles (
  id   TEXT PRIMARY KEY,
  --
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  content TEXT NOT NULL,
  --
  author TEXT NOT NULL,
  published BOOLEAN NOT NULL,
  --
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id TEXT PRIMARY KEY,
    --
    article_id TEXT NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    author_name TEXT NOT NULL,
    content TEXT NOT NULL,
    approved BOOLEAN NOT NULL DEFAULT FALSE,
    --
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
