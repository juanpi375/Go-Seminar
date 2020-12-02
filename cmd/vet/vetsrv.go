// This class is the main. Creates the DB, sets
// the configuration and starts the service

package main


import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	// "time" // Uncomment for inserting one animal
	"github.com/juanpi375/Go-Seminary/internal/config"
	"github.com/juanpi375/Go-Seminary/internal/database"
	"github.com/juanpi375/Go-Seminary/internal/service/vet"
)



func main(){
	cfg := readConfig()
	
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


// Reads the configuration of the config file
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


// Creates the schema when it doesn't exist
func createSchema (db *sqlx.DB) error{
	schema := `CREATE TABLE IF NOT EXISTS animals (
		id integer primary key autoincrement,
		name varchar,
		age integer);`
	
	_, err := db.Exec(schema)
	if err != nil{
		return err
	}

	// Uncomment only for inserting one extra animal!
	
	// insertAnimal := `INSERT INTO animals (name, age) VALUES (?,?)`
	// name := fmt.Sprintf("Name of %v", time.Now().Nanosecond())
	// age := time.Now().Nanosecond()
	// db.MustExec(insertAnimal, name, age)
	return nil
}