package user

import (
	"github.com/isuraem/hex/internal/models"
	"github.com/isuraem/hex/internal/ports"
)

type UserService struct {
	db ports.UserService
}

func NewUserService(db ports.UserService) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) AddUser(user models.User) error {
	return s.db.AddUser(user)
}

func (s *UserService) RemoveUser(id int) error {
	return s.db.RemoveUser(id)
}

func (s *UserService) ViewUser(id int) (models.User, error) {
	return s.db.ViewUser(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.db.ListUsers()
}
