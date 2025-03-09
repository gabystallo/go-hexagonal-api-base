package domain

//go:generate mockgen -destination=../mocks/domain/mockPersonRepository.go -package=domain . PersonRepository
type PersonRepository interface {
	GetAll() ([]Person, error)
	GetById(int) (*Person, error)
	Create(*Person) error
}
