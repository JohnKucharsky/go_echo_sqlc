package handler

import (
	"context"
	"github.com/JohnKucharsky/go_echo_sqlc/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) ProductPost(c echo.Context) error {
	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := apiConfig.Database.DB.CreateProduct(
		context.Background(),
		serializer.ProductBodyToProductCreate(productBody),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		serializer.DatabaseProductToProduct(product),
	)
}

func (apiConfig *DatabaseController) GetProducts(c echo.Context) error {
	products, err := apiConfig.Database.DB.GetProducts(
		context.Background(),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		serializer.DatabaseProductsToProducts(products),
	)
}

func (apiConfig *DatabaseController) GetOneProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	product, err := apiConfig.Database.DB.GetOneProduct(
		context.Background(),
		dbId,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusOK, serializer.DatabaseProductToProduct(product))
}

func (apiConfig *DatabaseController) UpdateProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := apiConfig.Database.DB.UpdateProduct(
		context.Background(),
		serializer.ProductBodyToProductUpdate(productBody, dbId),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		serializer.DatabaseProductToProduct(product),
	)
}

func (apiConfig *DatabaseController) DeleteProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	err = apiConfig.Database.DB.DeleteProduct(
		context.Background(),
		dbId,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
