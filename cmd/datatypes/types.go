package datatypes

type UserStore interface {
	CreateUser(Employee) (int64, error)
	GetUsers(*Employee) error
}

type Users struct {
	Id        uint8  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
