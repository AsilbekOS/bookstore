package connect

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	Host    = "localhost"
	Port    = 5432
	User    = "postgres"
	Passord = "7479"
	Dbname  = "bookstore_n9"
)

func OpenDb() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	data := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER_DB"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, err := sql.Open("postgres", data)
	if err != nil {
		log.Panic(err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping", err)
	}

	fmt.Println("Successfully connected database")

	return db
}
