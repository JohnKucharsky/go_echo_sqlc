package main

import (
	"github.com/JohnKucharsky/go_echo_sqlc/db"
	"github.com/JohnKucharsky/go_echo_sqlc/handler"
	"github.com/labstack/echo/v4"
)

func routes(route *echo.Group, dbConnectionString string) {
	database := db.DatabaseConnection(dbConnectionString)
	h := handler.DatabaseController{Database: database}

	route.GET("/healthz", handler.CheckHealth)

	// users
	route.POST("/users", h.UserPost)
	route.GET("/users", h.GetUsers)
	route.GET("/users/:id", h.GetOneUser)
	route.PUT("/users/:id", h.UpdateUser)
	route.DELETE("/users/:id", h.DeleteUser)
	// end users

	// products
	route.POST("/products", h.ProductPost)
	route.GET("/products", h.GetProducts)
	route.GET("/products/:id", h.GetOneProduct)
	route.PUT("/products/:id", h.UpdateProduct)
	route.DELETE("/products/:id", h.DeleteProduct)
	// end products
}
