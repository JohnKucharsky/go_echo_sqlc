package handler

import "github.com/JohnKucharsky/go_echo_sqlc/db"

type DatabaseController struct {
	Database *db.ApiConfig
}
