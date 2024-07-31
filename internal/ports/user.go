package ports

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService interface {
	AddUser(user User) error
	RemoveUser(id int) error
	ViewUser(id int) (User, error)
	ListUsers() ([]User, error)
}
