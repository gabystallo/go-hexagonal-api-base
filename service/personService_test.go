package service_test

import (
	"encoding/json"
	"testing"

	"github.com/gabystallo/go-hexagonal-api-base/domain"
	mockdomain "github.com/gabystallo/go-hexagonal-api-base/mocks/domain"
	"github.com/gabystallo/go-hexagonal-api-base/service"

	"github.com/golang/mock/gomock"
)

func TestCalling_GetAll_returns_correct_person_list(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	mr := mockdomain.NewMockPersonRepository(ctrl)
	mr.EXPECT().GetAll().Return([]domain.Person{
		{Id: 1, Name: "Gaby", City: "Ramos"},
		{Id: 2, Name: "Babu", City: "Castelar"},
	}, nil)
	s := service.NewPersonService(mr)

	// act
	persons, err := s.GetAll()

	// assert
	if err != nil {
		t.Errorf("Error calling GetAll: %s", err)
	}
	if len(persons) != 2 {
		t.Errorf("Wrong response count %+v", persons)
	}
	p, err := json.Marshal(persons[0])
	if err != nil {
		t.Errorf("Error marshalling person %+v", persons[0])
	}
	if string(p) != `{"id":1,"name":"Gaby","city":"Ramos"}` {
		t.Errorf("Person different than expected %s", p)
	}
}
