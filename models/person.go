package models

import "encoding/json"

type Person interface {
	GetName() string
	GetEmail() string
	GetAge() int
}

type PersonData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (p *PersonData) GetName() string {
	return p.Name
}

func (p *PersonData) GetEmail() string {
	return p.Email
}

func (p *PersonData) GetAge() int {
	return p.Age
}

func NewPerson(name, email string, age int) Person {
	return &PersonData{
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func PersonFromJSON(data []byte) (Person, error) {
	var p PersonData
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
