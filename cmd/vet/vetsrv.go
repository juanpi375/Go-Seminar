package main

import (
	"os"
	"fmt"
	"flag"
	"Go-Seminary/internal/config"
	"././internal/service/vet"
)

func main(){
	cfg := readConfig()

	// fmt.Println(cfg.Db.Driver)
	// fmt.Println(cfg.Version)

	service, err := vet.New(cfg)
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