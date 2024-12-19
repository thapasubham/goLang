package service

import (
	"database/sql"
	"fmt"

	"github.com/thapasubham/go-learn/cmd/datatypes"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}

}

func (s *Store) CreateUser(employee datatypes.Employee) (int64, error) {

	result, err := s.db.Exec("INSERT INTO employee (name, address) VALUES (?, ?)", employee.Name, employee.Address)
	if err != nil {
		return 0, err
	}
	fmt.Println(employee.Name + " " + employee.Address)

	id, err := result.LastInsertId()

	if id == 0 {
		fmt.Println("No rows affected. Possible constraint or query issue.")

	}

	fmt.Print(id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *Store) GetUser(user *datatypes.Employee) error {

	row, err := s.db.Query("Select name, address from employee LIMIT 5")

	if err == nil {
		for row.Next() {
			if err := row.Scan(&user.Name, &user.Address); err != nil {
				return err
			}
		}
	}
	return err
}

func (s *Store) EditUser(id int, user datatypes.Employee) error {

	result, err := s.db.Exec("Update employee set name =?, address =?  where id =?", user.Name, user.Address, id)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to fetch rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record with id %d may not exist", id)
	}

	return nil
}

func (s *Store) DeleteUser(id int) error {

	result, err := s.db.Exec("delete from employee where id =?", id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("Couldnt find the ")
	}
	if err != nil {
		return err
	}

	return nil
}
