package vet

import (
	"Go-Seminary/internal/config"
)

// Animal ...
type Animal struct{
	ID int64
	Name string 
	Age int
}

// VetService ...
type VetService interface{
	AddAnimal(Animal) error
	FindByID(int) *Animal
	FindAll() []*Animal
}

type service struct{
	conf *config.Config
}

// New ...
func New (c *config.Config) (VetService, error){
	return service{c}, nil
}

func (s service) AddAnimal(a Animal) error{ 
	return nil
}
func (s service) FindByID(id int) *Animal{
	return nil
}
func (s service) FindAll() []*Animal{
	var group []*Animal
	group = append(group, &Animal{0, "Some animal", 3})
	return group
}