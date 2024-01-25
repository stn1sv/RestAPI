package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testTask/models"
	"testTask/repository"
)

type PersonService struct {
	repo repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func getAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("getAge: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("getAge: %w", err)
	}

	age, ok := result["age"].(float64)
	if !ok {
		return 0, fmt.Errorf("age not found in the response")
	}

	return int(age), nil
}

func getGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("getGender: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("getGender: %w", err)
	}

	gender, ok := result["gender"].(string)
	if !ok {
		return "", fmt.Errorf("gender not found in the response")
	}

	return gender, nil
}

func getNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("getNationality: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("getNationality: %w", err)
	}

	countries, ok := result["country"].([]interface{})
	if !ok || len(countries) == 0 {
		return "", fmt.Errorf("nationality not found in the response")
	}

	nationality, ok := countries[0].(map[string]interface{})["country_id"].(string)
	if !ok {
		return "", fmt.Errorf("nationality not found in the response")
	}

	return nationality, nil
}

func (s *PersonService) AddPerson(person models.Person) (int, error) {
	if person.Name == "" || person.Surname == "" {
		log.Println("AddPerson: first or last name not entered")
		return 0, fmt.Errorf("error: first or last name not entered")
	}

	var err error
	person.Age, err = getAge(person.Name)
	if err != nil {
		log.Println("AddPerson: " + err.Error())
		return 0, err
	}
	person.Gender, err = getGender(person.Name)
	if err != nil {
		log.Println("AddPerson: " + err.Error())
		return 0, err
	}
	person.Nationality, err = getNationality(person.Name)
	if err != nil {
		log.Println("AddPerson: " + err.Error())
		return 0, err
	}

	log.Printf("Person %s added successfully", person)
	return s.repo.CreatePerson(person)
}

func (s *PersonService) GetPerson(params models.Person) ([]models.Person, error) {

	persons, err := s.repo.GetPerson(params)
	if err != nil {
		log.Println("GetPerson: " + err.Error())
		return []models.Person{}, err
	}

	return persons, nil
}

func (s *PersonService) DeletePerson(id string) error {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Println("DeletePerson: " + err.Error())
		return err
	}

	err = s.repo.DeletePerson(id_int)
	if err != nil {
		log.Println("DeletePerson: " + err.Error())
		return err
	}

	log.Printf("Person with id %s delete successfully", id)
	return nil
}

func (s *PersonService) UpdatePerson() {
	//TODO implement me
	panic("implement me")
}
