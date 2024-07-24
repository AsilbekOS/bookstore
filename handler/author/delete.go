package author

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "bookstore/postgres/connect"
)

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	} 


	query1 := `DELETE FROM book WHERE book_id = $1`
	res1, err := db.OpenDb().Exec(query1, id)
	if err != nil {
		log.Panic("Ma'lumotni o'chirishda xatolik:", err)
	}
	query := `DELETE FROM author WHERE author_id = $1`
	res, err := db.OpenDb().Exec(query, id)
	if err != nil {
		log.Panic("Ma'lumotni o'chirishda xatolik:", err)
	}

	rows1, err := res1.RowsAffected()
	if err != nil {
		log.Panic("Rowsda xatolik:", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Panic("Rowsda xatolik:", err)
	}

	if rows1 == 0 {
		fmt.Println("Bunday ID mavjud emas...")
		http.Error(w, "Bunday ID mavjud emas...",404)
		return
	}
	if rows == 0 {
		fmt.Println("Bunday ID mavjud emas...")
		http.Error(w, "Bunday ID mavjud emas...",404)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "%d - IDli Book o'chirildi", id)
	fmt.Fprintf(w,"\n")
	fmt.Fprintf(w, "%d - IDli Author o'chirildi", id)
}
