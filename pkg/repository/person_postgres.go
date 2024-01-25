package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"testTask/models"
)

type PersonPostgres struct {
	db *sqlx.DB
}

func NewPersonPostgres(db *sqlx.DB) *PersonPostgres {
	return &PersonPostgres{
		db: db,
	}
}

func (r *PersonPostgres) CreatePerson(person models.Person) (int, error) {
	var id int

	query := "INSERT INTO people (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRow(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PersonPostgres) GetPerson(params models.Person) ([]models.Person, error) {
	var persons []models.Person
	var conditions []string

	if params.Name != "" {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "name", params.Name))
	}
	if params.Surname != "" {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "surname", params.Surname))
	}
	if params.Patronymic != "" {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "patronymic", params.Patronymic))
	}
	if params.Age != 0 {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "age", strconv.Itoa(params.Age)))
	}
	if params.Gender != "" {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "gender", params.Gender))
	}
	if params.Nationality != "" {
		conditions = append(conditions, fmt.Sprintf("%s = '%v'", "nationality", params.Nationality))
	}
	where := strings.Join(conditions, " AND ")

	query := fmt.Sprintf("SELECT name, surname, patronymic, age, gender, nationality FROM people WHERE %s", where)
	err := r.db.Select(&persons, query)

	return persons, err
}

func (r *PersonPostgres) DeletePerson(id int) error {
	query := "DELETE FROM people WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonPostgres) UpdatePerson() error {
	//TODO implement me
	panic("implement me")
}
