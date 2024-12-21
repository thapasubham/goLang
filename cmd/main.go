package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/thapasubham/go-learn/cmd/api"
	"github.com/thapasubham/go-learn/cmd/utils"
)

func main() {

	dbURL := utils.LoadEnv("DB_URL")
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
