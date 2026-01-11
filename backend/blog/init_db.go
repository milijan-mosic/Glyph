package db

import (
	"context"
	"glyph/utils"
	"log"

	"github.com/jackc/pgx/v5"
)

var schema = `
	CREATE TABLE articles (
	  id   TEXT PRIMARY KEY,
	  --
	  title TEXT NOT NULL,
	  description  TEXT,
	  content TEXT NOT NULL,
	  --
	  author TEXT NOT NULL,
	  published BOOLEAN NOT NULL,
	  --
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

func InitializeDatabase() {
	ctx := context.Background()

	url := utils.GetDatabaseUrl()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s", err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, schema)
	if err != nil {
		log.Println("Failed to create table:", err)
	} else {
		log.Println("Table initialized successfully")
	}
}
