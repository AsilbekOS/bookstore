package book

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	bk "bookstore/handler/book/structs"
	db "bookstore/postgres/connect"
)	

func CreateBookMore(w http.ResponseWriter, r *http.Request) {
	b := []bk.Book{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Panic("Bookni create qilishda xatolik:", err)
	}

	for i := range b {
		bd, err := time.Parse("2006-01-02", b[i].Publication_date)
		if err != nil {
			log.Panic("Publicationni parse qilishda xatolik:", err)
		}

		b[i].Publication_date = bd.Format("2006-01-02")
		b[i].Created_at = time.Now()
		b[i].Updated_at = time.Now()

		database := db.OpenDb()
		defer database.Close()

		_, err = database.Exec("INSERT INTO book (title,author_id, publication_date, isbn, b_description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6,$7)", b[i].Title,b[i].Auther_id, b[i].Publication_date, b[i].Isbn, b[i].B_description, b[i].Created_at, b[i].Updated_at)
		if err != nil {
			log.Panic("Bookni tablega qo'shishda xatolik:", err)
		}
		
		w.WriteHeader(200)
		fmt.Println("Book tablega ma'lumot muvafaqqiyatli qo'shildi.")
	}
}

func CreateBookSin(w http.ResponseWriter, r *http.Request) {
	b := bk.Book{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Panic("Bookni create qilishda xatolik:", err)
	}
	bd, err := time.Parse("2006-01-02", b.Publication_date)
	if err != nil {
		log.Panic("Publicationni parse qilishda xatolik:", err)
	}

	b.Publication_date = bd.Format("2006-01-02")
	b.Created_at = time.Now()
	b.Updated_at = time.Now()

	database := db.OpenDb()
	defer database.Close()

	_, err = database.Exec("INSERT INTO book (title,author_id, publication_date, isbn, b_description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6,$7)", b.Title, b.Auther_id, b.Publication_date, b.Isbn, b.B_description, b.Created_at, b.Updated_at)
	if err != nil {
		log.Panic("Bookni tablega qo'shishda xatolik:", err)
	}

	w.WriteHeader(200)
	fmt.Println("Book tablega ma'lumot muvafaqqiyatli qo'shildi.")

}
