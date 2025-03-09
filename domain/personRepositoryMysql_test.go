package domain_test

import (
	"encoding/json"
	"hexapi/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestGetAll_calls_correct_query_and_returns_persons(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error initializing sqlmock")
	}
	rows := sqlmock.NewRows([]string{"id", "name", "city"}).
		AddRow(1, "Gaby", "Ramos").
		AddRow(2, "Babu", "Castelar")

	mock.ExpectQuery("from persons").WillReturnRows(rows)

	r := domain.NewPersonRepositoryMysql(sqlx.NewDb(db, "mysql"))

	// act
	persons, err := r.GetAll()

	// assert
	if err != nil {
		t.Errorf("Error calling GetAll %s", err)
	}
	if len(persons) != 2 {
		t.Errorf("Wrong amount of persons in response %+v", persons)
	}
	j, err := json.Marshal(persons[0].ToDto())
	if err != nil {
		t.Errorf("Error marshalling person %+v", persons[0])
	}
	if string(j) != `{"id":1,"name":"Gaby","city":"Ramos"}` {
		t.Errorf("Person different than expected %s", j)
	}
}
