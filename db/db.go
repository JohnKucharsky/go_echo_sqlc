package db

import (
	"database/sql"
	"github.com/JohnKucharsky/go_echo_sqlc/internal/database"
	"log"
)

type ApiConfig struct {
	DB *database.Queries
}

func DatabaseConnection(dbAddressString string) *ApiConfig {
	conn, err := sql.Open("postgres", dbAddressString)
	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}
	log.Print("Connected to db")
	db := database.New(conn)
	apiCfg := ApiConfig{DB: db}

	return &apiCfg
}
