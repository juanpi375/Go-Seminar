// This is the entities class
// Here the attributes and methods of
// the resources are set

package vet

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanpi375/Go-Seminary/internal/config"
)

// Animal ...
// Definition of attributes
type Animal struct{
	ID int64
	Name string 
	Age int
}

// Service ...
// Declaration of methods
type Service interface{
	FindByID(int) *Animal
	FindAll() []*Animal
	AddAnimal(Animal)
	DeleteAnimal(int)
	ReplaceAnimal(int, Animal)
}


// With the "New" method makes 
// a factory pattern to expose
// the functionality
type service struct{
	db *sqlx.DB
	conf *config.Config
}
// New ...
func New (db *sqlx.DB, c *config.Config) (Service, error){
	return service{db, c}, nil
}


// Definition of methods / Implementation of the interface------------
func (s service) FindAll() []*Animal{
	var group []*Animal
	if err := s.db.Select(&group, "SELECT * FROM animals"); err != nil{
		panic(err)
	}
	return group
}
func (s service) FindByID(id int) *Animal{
	var individual []*Animal
	query := `SELECT * FROM animals WHERE id = ?`
	if err := s.db.Select(&individual, query, id); err != nil{
		panic(err)
	}
	// Return the first as the the query returns an array with 1
	return individual[0]
}
func (s service) AddAnimal(a Animal){ 
	query := `INSERT INTO animals (name, age) VALUES (?,?)`
	_, err := s.db.Exec(query, a.Name, a.Age)
	if err != nil{
		panic(err)
	}
}
func (s service) DeleteAnimal(id int){
	query := `DELETE FROM animals WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil{
		panic(err)
	}
}
func (s service) ReplaceAnimal(id int, a Animal){ 
	query := `UPDATE animals SET name=?, age=? where id=?`
	_, err := s.db.Exec(query, a.Name, a.Age, id)
	if err != nil{
		panic(err)
	}
}
-----------------------------------------------------------------