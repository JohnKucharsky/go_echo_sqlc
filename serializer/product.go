package serializer

import (
	"github.com/JohnKucharsky/go_echo_sqlc/internal/database"
	"github.com/JohnKucharsky/go_echo_sqlc/utils"
	"time"
)

type ProductBody struct {
	Name   string  `json:"name" validate:"required"`
	Serial *string `json:"serial"`
}

type Product struct {
	ID        int32     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	ProductBody
}

func DatabaseProductToProduct(dbUser database.Product) Product {
	return Product{
		ID:        dbUser.ID,
		UpdatedAt: dbUser.UpdatedAt,
		ProductBody: ProductBody{
			Name:   dbUser.Name,
			Serial: utils.SqlNullStringToString(dbUser.Serial),
		},
	}
}

func ProductBodyToProductCreate(body ProductBody) database.CreateProductParams {
	return database.CreateProductParams{
		Name:      body.Name,
		Serial:    utils.StringToSqlNullString(body.Serial),
		UpdatedAt: time.Now().Local(),
	}
}

func ProductBodyToProductUpdate(
	body ProductBody,
	id int32,
) database.UpdateProductParams {
	return database.UpdateProductParams{
		Name:      body.Name,
		Serial:    utils.StringToSqlNullString(body.Serial),
		UpdatedAt: time.Now().Local(),
		ID:        id,
	}
}

func DatabaseProductsToProducts(dbProducts []database.Product) []Product {
	var products []Product

	for _, product := range dbProducts {
		products = append(products, DatabaseProductToProduct(product))
	}

	return products
}
