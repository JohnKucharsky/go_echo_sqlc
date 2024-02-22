package handler

import (
	"context"
	"github.com/JohnKucharsky/go_echo_sqlc/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) UserPost(c echo.Context) error {
	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := apiConfig.Database.DB.CreateUser(
		context.Background(),
		serializer.UserBodyToUserCreate(userBody),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, serializer.DatabaseUserToUser(user))
}

func (apiConfig *DatabaseController) GetUsers(c echo.Context) error {

	users, err := apiConfig.Database.DB.GetUsers(
		context.Background(),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusOK, serializer.DatabaseUsersToUsers(users))
}

func (apiConfig *DatabaseController) GetOneUser(c echo.Context) error {
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

	user, err := apiConfig.Database.DB.GetOneUser(
		context.Background(),
		dbId,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusOK, serializer.DatabaseUserToUser(user))
}

func (apiConfig *DatabaseController) UpdateUser(c echo.Context) error {
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

	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := apiConfig.Database.DB.UpdateUser(
		context.Background(),
		serializer.UserBodyToUserUpdate(userBody, dbId),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, serializer.DatabaseUserToUser(user))
}

func (apiConfig *DatabaseController) DeleteUser(c echo.Context) error {
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

	err = apiConfig.Database.DB.DeleteUser(
		context.Background(),
		dbId,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
