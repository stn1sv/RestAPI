package service

import (
	"testTask/models"
	"testTask/pkg/repository"
)

type Person interface {
	AddPerson(person models.Person) (int, error)
	GetPerson(params models.Person) ([]models.Person, error)
	DeletePerson(id string) error
	UpdatePerson(id string, params models.Person) error
}

type Service struct {
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewPersonService(repos),
	}
}
