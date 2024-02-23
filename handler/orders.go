package handler

import (
	"context"
	"github.com/JohnKucharsky/go_echo_sqlc/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) OrderPost(c echo.Context) error {
	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := apiConfig.Database.DB.CreateOrder(
		context.Background(),
		serializer.OrderBodyToOrderCreate(orderBody),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		product,
	)
}

func (apiConfig *DatabaseController) GetOrders(c echo.Context) error {
	orders, err := apiConfig.Database.DB.GetOrders(
		context.Background(),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		serializer.DatabaseOrdersToOrders(orders),
	)
}

func (apiConfig *DatabaseController) UpdateOrder(c echo.Context) error {
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

	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := apiConfig.Database.DB.UpdateOrder(
		context.Background(),
		serializer.OrderBodyToOrderUpdate(orderBody, dbId),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		order,
	)
}

func (apiConfig *DatabaseController) DeleteOrder(c echo.Context) error {
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

	err = apiConfig.Database.DB.DeleteOrder(
		context.Background(),
		dbId,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
