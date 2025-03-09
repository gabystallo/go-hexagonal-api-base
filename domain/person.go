package domain

import "hexapi/dto"

type Person struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	City string `db:"city"`
}

func (p Person) ToDto() dto.PersonResponse {
	return dto.PersonResponse{
		Id:   p.Id,
		Name: p.Name,
		City: p.City,
	}
}
