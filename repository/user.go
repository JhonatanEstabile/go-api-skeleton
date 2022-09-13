package repository

import (
	"api/db"
	"database/sql"

	"github.com/oklog/ulid/v2"
)

type IUser interface {
	CreateUser(name string, email string) (sql.Result, error)
}

type user struct {
	db *sql.DB
}

func NewUserRepository() *user {
	return &user{
		db: db.GetClient(),
	}
}

func (user *user) CreateUser(name string, email string) (sql.Result, error) {
	return user.db.Exec(
		`
			INSERT INTO users (
				id,
				name,
				email
			)
			VALUES (
				?,
				?,
				?
			)
		`,
		ulid.Make().String(),
		name,
		email,
	)
}
