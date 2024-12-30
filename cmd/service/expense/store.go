package expense

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

// get all the expenses made by the user
func (s *Store) GetExpenses(expenses *[]datatypes.Expense, user_id int) error {

	result, err := s.db.Query("Select id, user_id, description, date, category, amount from expenses where user_id =?", user_id)

	if err != nil {
		return err
	}
	found := false
	defer result.Close()
	var expense datatypes.Expense
	for result.Next() {
		found = true
		err := result.Scan(&expense.ID, &expense.UserID, &expense.Description, &expense.Date, &expense.Category, &expense.Amount)

		if err != nil {
			return err
		}
		*expenses = append(*expenses, expense)
	}
	if !found {
		return fmt.Errorf("no expense found with user_id %v", user_id)
	}
	return nil
}

// get a single expense made by the user
func (s *Store) GetExpense(expense *datatypes.Expense, id int) error {

	result, err := s.db.Query("Select id, user_id, description, date, category, amount from expenses where id =?", id)

	if err != nil {
		return err
	}

	found := false
	for result.Next() {
		found = true
		err := result.Scan(&expense.ID, &expense.UserID, &expense.Description, &expense.Date, &expense.Category, &expense.Amount)

		if err != nil {
			return err
		}

	}
	if !found {
		return fmt.Errorf("no expense found")
	}
	return nil
}

// Add a new expense to the database
func (s *Store) AddExpense(expense *datatypes.Expense, userID int) error {
	result, err := s.db.Exec(
		"INSERT INTO expenses (user_id, category, amount, description, date) VALUES (?, ?, ?, ?, ?)",
		userID, expense.Category, expense.Amount, expense.Description, expense.Date,
	)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %v", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return nil
}

// Edit an existing expense in the database
func (s *Store) EditExpense(expense *datatypes.Expense, userID int) error {
	result, err := s.db.Exec(
		"UPDATE expenses SET category = ?, amount = ?, description = ?, date = ? WHERE id = ? AND user_id = ?",
		expense.Category, expense.Amount, expense.Description, expense.Date, expense.ID, userID,
	)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no expense found with the given id and user_id")
	}

	return nil
}

// delete an expense
func (s *Store) DeleteExpense(e_id int, user_id int) error {

	result, err := s.db.Exec("delete from expenses where id = ? and user_id =?", e_id, user_id)

	if err != nil {
		return fmt.Errorf("failed to delete %v", err)

	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no expense found with the given id and user_id")
	}
	return nil
}
