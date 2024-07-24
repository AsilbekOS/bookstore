package author

import (
	athr "bookstore/handler/author/structs"
	dtb "bookstore/postgres/connect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	db := dtb.OpenDb()
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var a athr.Author
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Authorni create qilishda xatolik:", err)
		return
	}

	bd, err := time.Parse("2006-01-02", a.Birth_date)
	if err != nil {
		log.Println("Birthdateni parse qilishda xatolik:", err)
	}

	a.Birth_date = bd.Format("2006-01-02")
	a.Updated_at = time.Now()

	query := `UPDATE author 
	    SET name = $1, birth_date = $2, biography = $3, updated_at = $4
    	WHERE author_id = $5
    	RETURNING author_id, name, birth_date, biography, created_at, updated_at`
	row := db.QueryRow(query, a.Name, a.Birth_date, a.Biography, a.Updated_at, id)

	var upAr athr.Author
	err = row.Scan(&upAr.Author_id, &upAr.Name, &upAr.Birth_date, &upAr.Biography, &upAr.Created_at, &upAr.Updated_at)
	if err != nil {
		http.Error(w, "Failed to update author", http.StatusInternalServerError)
		log.Println("Author tablega yangilashda xatolik:", err)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(upAr)

	fmt.Fprintf(w, "%d - IDli AUTHOR muvafaqqiyatli yangilandi.", id)
	fmt.Println("Author muvafaqqiyatli yangilandi.")
}
