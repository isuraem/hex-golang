package db

import (
	"database/sql"

	"github.com/isuraem/hex/internal/ports"
	_ "github.com/lib/pq"
)

type UserDB struct {
	conn *sql.DB
}

func NewUserDB(conn *sql.DB) *UserDB {
	return &UserDB{conn: conn}
}

func (db *UserDB) AddUser(user ports.User) error {
	_, err := db.conn.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (db *UserDB) RemoveUser(id int) error {
	_, err := db.conn.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (db *UserDB) ViewUser(id int) (ports.User, error) {
	row := db.conn.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	var user ports.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return ports.User{}, err
	}
	return user, nil
}

func (db *UserDB) ListUsers() ([]ports.User, error) {
	rows, err := db.conn.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []ports.User
	for rows.Next() {
		var user ports.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
