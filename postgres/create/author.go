package create

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTableAuthor(db *sql.DB) {
	createtableauthor := `CREATE TABLE IF NOT EXISTS author (
		author_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		birth_date DATE,
		biography TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL
	);`
	_, err := db.Exec(createtableauthor)
	if err != nil {
		log.Println("Table yaratishda xatolik:", err)
	}
	fmt.Println("Successfully created author table")
}
