package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // postgresql
)

// DB Connection
var DB *sql.DB

const (
	dbHost = "DB_HOST"
	dbPort = "DB_PORT"
	dbUser = "DB_USER"
	dbPass = "DB_PASS"
	dbName = "DB_NAME"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbHost)
	if !ok {
		panic("DB_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbPort)
	if !ok {
		panic("DB_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbUser)
	if !ok {
		panic("DB_USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbPass)
	if !ok {
		panic("DB_PASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbName)
	if !ok {
		panic("DB_NAME environment variable required but not set")
	}
	conf[dbHost] = host
	conf[dbPort] = port
	conf[dbUser] = user
	conf[dbPass] = password
	conf[dbName] = name
	return conf
}

// InitializeDB PostgreSQL
func InitializeDB() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbHost], config[dbPort],
		config[dbUser], config[dbPass], config[dbName])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established.")
}
