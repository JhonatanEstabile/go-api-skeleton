package repository

import (
	"api/db"
	"api/models"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
)

type IUser interface {
	Create(userData *models.User) (sql.Result, error)
	List(query string, args []interface{}) *[]models.User
	Detail(id string) *models.User
	Delete(id string) error
	Update(id string, userData *models.User) error
}

type user struct {
	db *sqlx.DB
}

func NewUserRepository() *user {
	return &user{
		db: db.GetClient(),
	}
}

func (user *user) Create(userData *models.User) (sql.Result, error) {
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
		userData.Name,
		userData.Email,
	)
}

func (user *user) List(query string, args []interface{}) *[]models.User {
	rows, err := user.db.Queryx(
		`SELECT * FROM users `+query,
		args...,
	)

	if err != nil {
		log.Fatalln(err)
	}

	var usersData []models.User

	for rows.Next() {
		userData := models.User{}
		rows.StructScan(&userData)
		usersData = append(usersData, userData)
	}

	return &usersData
}

func (user *user) Detail(id string) *models.User {
	userData := models.User{}

	err := user.db.QueryRowx(
		`
			SELECT * FROM users
			WHERE id = ?
			LIMIT 1
		`,
		id,
	).StructScan(&userData)

	if err != nil {
		log.Fatalln(err)
	}

	return &userData
}

func (user *user) Delete(id string) error {
	_, err := user.db.Exec(
		`
			DELETE FROM users
			WHERE id = ?
		`,
		id,
	)

	return err
}

func (user *user) Update(id string, userData *models.User) error {
	_, err := user.db.Exec(
		`
			UPDATE users
			SET
				name = ?,
				email = ?
			WHERE id = ?
		`,
		userData.Name,
		userData.Email,
		id,
	)

	return err
}
