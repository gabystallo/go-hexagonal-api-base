package handler

import (
	"encoding/json"
	"fmt"
	"hexapi/dto"
	"hexapi/service"
	"hexapi/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	service service.PersonService
}

func (h PersonHandler) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := h.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Error al obtener personas"))
		if err != nil {
			panic("Error writing to client")
		}
		return
	}
	utils.SendJSON(w, map[string][]dto.PersonResponse{"persons": persons})
}

func (h PersonHandler) GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al leer el ID"))
		return
	}
	person, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No se encontró la persona con ID: %v", id)))
		return
	}
	utils.SendJSON(w, map[string]dto.PersonResponse{"person": *person})
}

func (h PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	pr := dto.CreatePersonRequest{}
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf("Error al interpretar la persona")))
		return
	}
	err = h.service.Create(&pr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error al crear la persona")))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func NewPersonHandler(service service.PersonService) PersonHandler {
	return PersonHandler{service}
}
