package service

import (
	"hexapi/domain"
	"hexapi/dto"
)

//go:generate mockgen -destination=../mocks/service/mockPersonService.go -package=service . PersonService
type PersonService interface {
	GetAll() ([]dto.PersonResponse, error)
	GetById(int) (*dto.PersonResponse, error)
	Create(*dto.CreatePersonRequest) error
}

type DefaultPersonService struct {
	repo domain.PersonRepository
}

func (s DefaultPersonService) GetAll() ([]dto.PersonResponse, error) {
	persons, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.PersonResponse, len(persons))
	for i, p := range persons {
		response[i] = p.ToDto()
	}

	return response, nil
}

func (s DefaultPersonService) GetById(id int) (*dto.PersonResponse, error) {
	person, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	pr := person.ToDto()
	return &pr, err
}

func (s DefaultPersonService) Create(pr *dto.CreatePersonRequest) error {
	person := domain.Person{
		Name: pr.Name,
		City: pr.City,
	}
	err := s.repo.Create(&person)
	if err != nil {
		return err
	}
	return nil
}

func NewPersonService(r domain.PersonRepository) DefaultPersonService {
	return DefaultPersonService{r}
}
