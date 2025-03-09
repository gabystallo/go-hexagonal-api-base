package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gabystallo/go-hexagonal-api-base/dto"
	"github.com/gabystallo/go-hexagonal-api-base/handler"
	"github.com/gabystallo/go-hexagonal-api-base/mocks/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestGetAllPersons_returns_error_code_when_services_returns_error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	personServiceMock := service.NewMockPersonService(ctrl)
	personServiceMock.EXPECT().GetAll().Return(nil, errors.New("Some db error"))
	ph := handler.NewPersonHandler(personServiceMock)

	router := mux.NewRouter()
	router.HandleFunc("/persons", ph.GetAllPersons).Methods(http.MethodGet)

	request, _ := http.NewRequest(http.MethodGet, "/persons", nil)

	// act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// assert
	if recorder.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("Response code was not ok: %v", recorder.Result().StatusCode)
	}
	response, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	if string(response) != "Error al obtener personas" {
		t.Errorf("Unexpected response error message: %s", response)
	}

}

func TestGetAllPersons_returns_list_of_persons_with_200(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	personServiceMock := service.NewMockPersonService(ctrl)
	expectedServiceResponse := []dto.PersonResponse{
		{Id: 1, Name: "Gaby", City: "Ramos"},
		{Id: 2, Name: "Baque", City: "Morón"},
	}
	expectedResponse, _ := json.Marshal(map[string]any{"persons": expectedServiceResponse})
	personServiceMock.EXPECT().GetAll().Return(expectedServiceResponse, nil)
	ph := handler.NewPersonHandler(personServiceMock)

	router := mux.NewRouter()
	router.HandleFunc("/persons", ph.GetAllPersons).Methods(http.MethodGet)

	request, _ := http.NewRequest(http.MethodGet, "/persons", nil)

	// act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// assert
	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Response code was not ok: %v", recorder.Result().StatusCode)
	}
	response, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	if strings.TrimSpace(string(expectedResponse)) != strings.TrimSpace(string(response)) {
		t.Errorf("Response does not match expected response\nExpected: %s\nActual: %s", expectedResponse, response)
	}

}
