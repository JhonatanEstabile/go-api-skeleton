package user

import (
	"api/dependencies/repositories/mysql"
	domain "api/domain/user"
)

//type IUser interface {
//	Create(userData *user) (sql.Result, error)
//	List(query string, args []interface{}) *[]user2.User
//	Detail(id string) *user2.User
//	Delete(id string) error
//	Update(id string, userData *user2.User) error
//}

type Repository struct {
	*mysql.Repository
	//Ulid
}

func NewUserRepository(mysql *mysql.Repository) *Repository {
	return &Repository{
		Repository: mysql,
	}
}

func (r *Repository) Create(user *domain.User) (*domain.User, error) {
	return r.Client.Exec(
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
		user.Name,
		user.Email,
	)
}

func (r *Repository) Find(query string, args []interface{}) ([]*domain.User, error) {
	rows, err := r.Client.Queryx(
		`SELECT * FROM users `+query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for rows.Next() {
		u := domain.User{}
		rows.StructScan(&u)
		users = append(users, &u)
	}

	return users, nil
}

func (r *Repository) FindOne(id string) *domain.User {
	u := domain.User{}
	err := r.Client.QueryRowx(
		`
			SELECT * FROM users
			WHERE id = ?
			LIMIT 1
		`,
		id,
	).StructScan(&u)

	if err != nil {
		return nil
	}
	return &u
}

func (r *Repository) Delete(id string) error {
	_, err := r.Client.Exec(
		`
			DELETE FROM users
			WHERE id = ?
		`,
		id,
	)
	return err
}

func (r *Repository) Update(id string, user *domain.User) error {
	_, err := r.Client.Exec(
		`
			UPDATE users
			SET
				name = ?,
				email = ?
			WHERE id = ?
		`,
		user.Name,
		user.Email,
		id,
	)
	return err
}
