package db

import (
	"context"
	"log"

	"github.com/isuraem/hex/internal/adapters/framework/right/db"
	"github.com/isuraem/hex/internal/models"
)

type UserDB struct {
	adapter *db.Adapter
}

func NewUserDB(adapter *db.Adapter) *UserDB {
	return &UserDB{adapter: adapter}
}

func (db *UserDB) AddUser(user models.User) error {
	_, err := db.adapter.DB.NewInsert().
		Model(&user).
		Exec(context.Background())

	if err != nil {
		log.Printf("Error executing query: %v", err)
	}

	return err
}

func (db *UserDB) RemoveUser(id int) error {
	_, err := db.adapter.DB.NewDelete().
		Model((*models.User)(nil)).
		Where("id = ?", id).
		Exec(context.Background())

	return err
}

func (db *UserDB) ViewUser(id int) (models.User, error) {
	var user models.User
	err := db.adapter.DB.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(context.Background())

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (db *UserDB) ListUsers() ([]models.User, error) {
	var users []models.User
	err := db.adapter.DB.NewSelect().
		Model(&users).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return users, nil
}
