package domain

import "github.com/jmoiron/sqlx"

type PersonRepositoryMysql struct {
	db *sqlx.DB
}

func (r PersonRepositoryMysql) GetAll() ([]Person, error) {
	persons := []Person{}
	err := r.db.Select(&persons, "SELECT * from persons")
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (r PersonRepositoryMysql) GetById(id int) (*Person, error) {
	person := Person{}
	err := r.db.Get(&person, "SELECT * from persons where id = ?", id)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r PersonRepositoryMysql) Create(p *Person) error {
	_, err := r.db.Exec("INSERT INTO persons (name, city) values (?, ?)", p.Name, p.City)
	return err
}

func NewPersonRepositoryMysql(db *sqlx.DB) PersonRepositoryMysql {
	return PersonRepositoryMysql{db}
}
