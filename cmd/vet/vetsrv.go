package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	// "time" // Uncomment for inserting one animal
	// "Go-Seminary/internal/config"
	"github.com/juanpi375/Go-Seminary/internal/config"
	// "Go-Seminary/internal/service/vet"
	"github.com/juanpi375/Go-Seminary/internal/database"
	"github.com/juanpi375/Go-Seminary/internal/service/vet"
)

func main(){
	cfg := readConfig()
	// fmt.Println(cfg.Db.Driver)
	// fmt.Println(cfg.Version)
	
	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	service, _ := vet.New(db, cfg) 
	httpService := vet.NewHTTPTransport(service)
	router := gin.Default()
	httpService.Register(router)
	router.Run()

	for _, elem := range service.FindAll(){
		fmt.Println(elem)
	}
}

func readConfig() *config.Config{
	configFile := flag.String("config", "./config.yaml", "This is the config service")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
} 

// Creates the schema if not exists
func createSchema (db *sqlx.DB) error{
	schema := `CREATE TABLE IF NOT EXISTS animals (
		id integer primary key autoincrement,
		name varchar,
		age integer);`
	
	// Execute a query on the server
	_, err := db.Exec(schema)
	if err != nil{
		return err
	}

	// Uncomment only for inserting one extra animal!
	// // or, you can use MustExec, which panics on error
	// insertAnimal := `INSERT INTO animals (name, age) VALUES (?,?)`
	// name := fmt.Sprintf("Name of %v", time.Now().Nanosecond())
	// age := time.Now().Nanosecond()
	// db.MustExec(insertAnimal, name, age)
	return nil
}