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
