package ports

import "github.com/isuraem/hex/internal/models"

type UserService interface {
	AddUser(user models.User) error
	RemoveUser(id int) error
	ViewUser(id int) (models.User, error)
	ListUsers() ([]models.User, error)
}
