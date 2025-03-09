package domain

type PersonRepositoryStub struct {
	stubbedData []Person
}

func (r PersonRepositoryStub) GetAll() ([]Person, error) {
	return r.stubbedData, nil
}

func (r PersonRepositoryStub) GetById() (*Person, error) {
	return &r.stubbedData[0], nil
}

func (r PersonRepositoryStub) Create(p *Person) error {
	return nil
}

func NewPersonRepositoryStub() PersonRepositoryStub {
	return PersonRepositoryStub{stubbedData: []Person{
		{Id: 1, Name: "Gaby", City: "Ramos"},
		{Id: 2, Name: "Babu", City: "Castelar"},
	}}
}
