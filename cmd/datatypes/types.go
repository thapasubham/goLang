package datatypes

type UserStore interface {
	CreateUser(User) (int64, error)
	LogIn(User) (int, error)
}

type ExpenseStore interface {
	GetExpenses(*[]Expense, int) error
	GetExpense(*Expense, int) error
	AddExpense(*Expense, int) error
	EditExpense(*Expense, int) error
	DeleteExpense(int, int) error
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Expense struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}
