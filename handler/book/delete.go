package book

import (
	db "bookstore/postgres/connect"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM book WHERE book_id = $1`
	res, err := db.OpenDb().Exec(query, id)
	if err != nil {
		log.Panic("Ma'lumotni o'chirishda xatolik:", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Panic("Rowsda xatolik:", err)
	}

	if rows == 0 {
		fmt.Println("Bunday ID mavjud emas...")
		http.Error(w, "Bunday ID mavjud emas...", 404)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "%d - IDli Book o'chirildi", id)
}
