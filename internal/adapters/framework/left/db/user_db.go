package db

import (
	"log"

	"github.com/isuraem/hex/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/isuraem/hex/internal/adapters/framework/right/db"
)

type UserDB struct {
	adapter *db.Adapter
}

func NewUserDB(adapter *db.Adapter) *UserDB {
	return &UserDB{adapter: adapter}
}

func (db *UserDB) AddUser(user models.User) error {
	queryString, args, err := sq.Insert("users").Columns("name", "email").
		Values(user.Name, user.Email).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	// Log the query and args for debugging
	log.Printf("Executing query: %s with args: %v", queryString, args)

	_, err = db.adapter.DB.Exec(queryString, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
	}

	return err
}

func (db *UserDB) RemoveUser(id int) error {
	queryString, args, err := sq.Delete("users").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = db.adapter.DB.Exec(queryString, args...)
	return err
}

func (db *UserDB) ViewUser(id int) (models.User, error) {
	queryString, args, err := sq.Select("id", "name", "email").From("users").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	row := db.adapter.DB.QueryRow(queryString, args...)
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (db *UserDB) ListUsers() ([]models.User, error) {
	queryString, args, err := sq.Select("id", "name", "email").From("users").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.adapter.DB.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
