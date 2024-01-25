package repository

import (
	"github.com/jmoiron/sqlx"
	"testTask/models"
)

type Person interface {
	CreatePerson(person models.Person) (int, error)
	GetPerson(params models.Person) ([]models.Person, error)
	DeletePerson(id int) error
	UpdatePerson(id int, params models.Person) error
}

type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewPersonPostgres(db),
	}
}
