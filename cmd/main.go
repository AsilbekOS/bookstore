package main

import (
	dta "bookstore/handler/author"
	dtb "bookstore/handler/book"
	// dt "bookstore/postgres/connect"
	// crt "bookstore/postgres/create"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// db := dt.OpenDb()
	// crt.CreateTableAuthor(db)
	// crt.CreateTableBook(db)


	http.HandleFunc("POST /createall", dta.CreateAuthorMore)
	http.HandleFunc("POST /createone", dta.CreateAuthorSin)
	http.HandleFunc("DELETE /author/{id}", dta.DeleteAuthor)
	http.HandleFunc("GET /readid/{id}", dta.ReadwithID)
	http.HandleFunc("GET /readall", dta.ReadAll)
	http.HandleFunc("PUT /update/{id}", dta.UpdateAuthor)

	http.HandleFunc("POST /createbookmore", dtb.CreateBookMore)
	http.HandleFunc("POST /createbookone", dtb.CreateBookSin)
	http.HandleFunc("DELETE /book/{id}", dtb.DeleteBook)
	http.HandleFunc("GET /readbookid/{id}", dtb.ReadwithID)
	http.HandleFunc("GET /readbookall", dtb.ReadAll)
	http.HandleFunc("PUT /updatebook/{id}", dtb.UpdateBook)


	fmt.Println("Server is listening: 8080")
	http.ListenAndServe(":8080", nil)
}
