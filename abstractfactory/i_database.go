package main

import (
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	SetParameters(host string, port int, user string, password string, database string)
	SetUpConnection() (error)
	ExecuteQuery(query string) (interface{}, error)
}

type Database struct {
	host    string
	port    int
	user    string
	password string
	database string

	db *gorm.DB
}

func (db *Database) SetParameters(host string, port int, user string, password string, database string) {
	db.host = host
	db.port = port
	db.user = user
	db.password = password
	db.database = database
} 

func (db *Database) SetUpConnection() (error) {
	dsn := "host=" + db.host + " port=" + strconv.Itoa(db.port) + " user=" + db.user + " password=" + db.password + " dbname=" + db.database + " sslmode=disable"
	var err error
	db.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	return err
}

func (db *Database) ExecuteQuery(query string) (any, error) {
	// Execute the query using GORM
	var result []map[string]any
	err := db.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}