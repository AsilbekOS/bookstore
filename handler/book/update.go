package book

import (
	bk "bookstore/handler/book/structs"
	dtb "bookstore/postgres/connect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := dtb.OpenDb()
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var b bk.Book
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Bookni create qilishda xatolik:", err)
		return
	}

	bd, err := time.Parse("2006-01-02", b.Publication_date)
	if err != nil {
		log.Println("publication_date parse qilishda xatolik:", err)
	}

	b.Publication_date = bd.Format("2006-01-02")
	b.Updated_at = time.Now()

	query := `UPDATE book
	    SET title = $1, publication_date = $2, isbn = $3, b_description = $4, updated_at = $5
    	WHERE book_id = $6
    	RETURNING book_id, title, author_id, publication_date, isbn, b_description, created_at, updated_at`
	row := db.QueryRow(query, b.Title, b.Publication_date, b.Isbn, b.B_description, b.Updated_at, id)

	var upBk bk.Book
	err = row.Scan(&upBk.Book_id, &upBk.Title, &upBk.Auther_id, &upBk.Publication_date, &upBk.Isbn, &upBk.B_description, &upBk.Created_at, &upBk.Updated_at)
	if err != nil {
		http.Error(w, "Book yangilanishda xatolik... ", http.StatusInternalServerError)
		log.Println("Book yangilashda xatolik... ", err)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(upBk)
	
	json.NewEncoder(w).Encode("Book muvafaqqiyatli yangilandi.")
	fmt.Println("Book muvafaqqiyatli yangilandi.")
}
