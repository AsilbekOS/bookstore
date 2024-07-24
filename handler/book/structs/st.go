package structs

import (
	"time"
)

type Book struct {
	Book_id          int       `json:"book_id"`
	Title            string    `json:"title"`
	Auther_id        int       `json:"author_id"`
	Publication_date string    `json:"publication_date"`
	Isbn             string    `json:"isbn"`
	B_description    string    `json:"b_description"`
	Created_at       time.Time `json:"create_at"`
	Updated_at       time.Time `json:"updated_at"`
}
