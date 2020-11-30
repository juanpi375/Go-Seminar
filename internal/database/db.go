package database

import (
	"github.com/juanpi375/Go-Seminary/internal/config"
	"github.com/jmoiron/sqlx"
	"errors"
	_ "github.com/mattn/go-sqlite3" // Support for sqlite driver
)

// NewDatabase ..
func NewDatabase (conf *config.Config) (*sqlx.DB, error){
	switch conf.Db.Type{
		case "sqlite3":
			db, err := sqlx.Open(conf.Db.Driver, conf.Db.Conn)
			if err != nil{
				return nil, err
			}

			err = db.Ping()
			if err != nil{
				return nil, err
			}

			return db, nil
		default:
			return nil, errors.New("Invalid DB, please check Type")
	}
}
