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
