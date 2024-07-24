package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	bk "bookstore/handler/book/structs"
	dtb "bookstore/postgres/connect"
)

func ReadwithID(w http.ResponseWriter, r *http.Request) {
	db := dtb.OpenDb()
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM book WHERE book_id = $1"
	row, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}
	var books []bk.Book
	for row.Next() {
		var book bk.Book
		err = row.Scan(&book.Book_id, &book.Title, &book.Auther_id, &book.Publication_date, &book.Isbn, &book.B_description, &book.Created_at, &book.Updated_at)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}

	fmt.Println(books)
	fmt.Println("Ma'lumot o'qib olindi...")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	db := dtb.OpenDb()
	query := "SELECT * FROM book"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var books []bk.Book
	for row.Next() {
		var book bk.Book
		err = row.Scan(&book.Book_id, &book.Title, &book.Auther_id, &book.Publication_date, &book.Isbn, &book.B_description, &book.Created_at, &book.Updated_at)		
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}

	fmt.Println("Ma'lumot o'qib olindi...")

	fmt.Println(books)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}
