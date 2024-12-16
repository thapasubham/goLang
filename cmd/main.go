package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/thapasubham/go-learn/cmd/api"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env" + err.Error())
	}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbURL)
	initDatabase(db)
	if err != nil {
		panic(err)
	}
	server := api.NewApiServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDatabase(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!!")
}
