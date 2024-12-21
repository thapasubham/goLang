package service

import (
	"database/sql"
	"fmt"

	"github.com/thapasubham/go-learn/cmd/datatypes"
	"github.com/thapasubham/go-learn/cmd/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}

}

func (s *Store) CreateUser(user datatypes.User) (int64, error) {

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	var existingUser string
	err = s.db.QueryRow("SELECT username FROM users WHERE username = ?", user.Username).Scan(&existingUser)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if existingUser != "" {
		return 0, fmt.Errorf("username already exists")
	}

	result, err := s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (s *Store) LogIn(user datatypes.User) (int, error) {
	var storedPasswordHash string
	var id int
	err := s.db.QueryRow("SELECT id, password FROM users WHERE username = ?", user.Username).Scan(&id, &storedPasswordHash)
	if err != nil {
		return 0, fmt.Errorf("username not found")
	}

	if !utils.CheckPasswordHash(user.Password, storedPasswordHash) {
		return 0, fmt.Errorf("incorrect password")
	}

	return id, nil
}
