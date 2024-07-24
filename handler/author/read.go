package author

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	athr "bookstore/handler/author/structs"
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

	query := "SELECT * FROM author WHERE author_id = $1"
	row, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}
	var authors []athr.Author
	for row.Next() {
		var author athr.Author
		err = row.Scan(&author.Author_id, &author.Name, &author.Birth_date, &author.Biography, &author.Created_at, &author.Updated_at)
		if err != nil {
			fmt.Println(err)
		}
		authors = append(authors, author)
	}

	fmt.Println(authors)
	fmt.Println("Ma'lumot o'qib olindi...")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	db := dtb.OpenDb()
	query := "SELECT * FROM author"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var authors []athr.Author
	for row.Next() {
		var author athr.Author
		err = row.Scan(&author.Author_id, &author.Name, &author.Birth_date, &author.Biography, &author.Created_at, &author.Updated_at)
		if err != nil {
			fmt.Println(err)
		}
		authors = append(authors, author)
	}
	
	fmt.Println("Ma'lumot o'qib olindi...")
	
	fmt.Println(authors)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authors)
}