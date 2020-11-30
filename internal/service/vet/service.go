package vet

import (
	"github.com/juanpi375/Go-Seminary/internal/config"
	"github.com/jmoiron/sqlx"
)

// Animal ...
type Animal struct{
	ID int64
	Name string 
	Age int
}

// Service ...
type Service interface{
	AddAnimal(Animal) error
	FindByID(int) *Animal
	FindAll() []*Animal
}

type service struct{
	db *sqlx.DB
	conf *config.Config
}

// New ...
func New (db *sqlx.DB, c *config.Config) (Service, error){
	return service{db, c}, nil
}

func (s service) AddAnimal(a Animal) error{ 
	return nil
}
func (s service) FindByID(id int) *Animal{
	return nil
}
func (s service) FindAll() []*Animal{
	var group []*Animal
	if err := s.db.Select(&group, "SELECT * FROM animals"); err != nil{
		panic(err)
	}
	return group
}