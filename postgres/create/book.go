package create

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTableBook(db *sql.DB) {
	createtablebook := `CREATE TABLE IF NOT EXISTS book (
		book_id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		authors_id INT,
		publication_date DATE,
		isbn VARCHAR(255),
		b_description TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL,
		FOREIGN KEY (authors_id) REFERENCES author (author_id)
	);`

	_, err := db.Exec(createtablebook)
	if err != nil {
		log.Fatal("Book tableni yaratishda xatolik:", err)
	}

	fmt.Println("Successfully created book table")
}
