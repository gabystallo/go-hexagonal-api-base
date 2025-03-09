package dto

type PersonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type CreatePersonRequest struct {
	Name string `json:"name"`
	City string `json:"city"`
}
