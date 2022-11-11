package models

type User struct {
	ID    string `db:"id"`
	Name  string `json:"name" validate:"required,min=3,max=32" db:"name"`
	Email string `json:"email" validate:"required,email" db:"email"`
}

var UserFields []string = []string{
	"Name",
	"Email",
}
