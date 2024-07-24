package author

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	athr "bookstore/handler/author/structs"
	db "bookstore/postgres/connect"
)

// type Author struct {
// 	Author_id  uuid.UUID `json:"authorid"`
// 	First_name string    `json:"firstname"`
// 	Bith_date  string    `json:"birthdate"`
// 	Biography  string    `json:"biography"`
// 	Created_at time.Time `json:"create_at"`
// 	Updated_at time.Time `json:"updated_at"`
// }

func CreateAuthorMore(w http.ResponseWriter, r *http.Request) {
	a := []athr.Author{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Panic("Authorni create qilishda xatolik:", err)
	}
	// fmt.Println(a)

	for i := range a {
		bd, err := time.Parse("2006-01-02", a[i].Birth_date)
		if err != nil {
			log.Panic("Birthdateni parse qilishda xatolik:", err)
		}

		a[i].Birth_date = bd.Format("2006-01-02")
		a[i].Created_at = time.Now()
		a[i].Updated_at = time.Now()

		w.WriteHeader(200)
		_, err = db.OpenDb().Exec("INSERT INTO author (name, birth_date, biography, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", a[i].Name, a[i].Birth_date, a[i].Biography, a[i].Created_at, a[i].Updated_at)
		if err != nil {
			log.Panic("Author tablega qo'shishda xatolik:", err)
		}
		fmt.Println("Author tablega ma'lumot muvafaqqiyatli qo'shildi.")
	}
}

func CreateAuthorSin(w http.ResponseWriter, r *http.Request) {
	a := athr.Author{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Panic("Authorni create qilishda xatolik:", err)
	}

	bd, err := time.Parse("2006-01-02", a.Birth_date)
	if err != nil {
		log.Panic("Birthdateni parse qilishda xatolik:", err)
	}

	a.Birth_date = bd.Format("2006-01-02")
	a.Created_at = time.Now()
	a.Updated_at = time.Now()

	w.WriteHeader(200)

	_, err = db.OpenDb().Exec("INSERT INTO author (name, birth_date, biography, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", a.Name, a.Birth_date, a.Biography, a.Created_at, a.Updated_at)
	if err != nil {
		log.Panic("Author tablega qo'shishda xatolik:", err)
	}
	fmt.Println("Author tablega ma'lumot muvafaqqiyatli qo'shildi.")

	
}
