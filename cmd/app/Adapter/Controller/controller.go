package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// driverで定義されたエンドポイントの関数を定義する

func GetUsers(c echo.Context) error {
	u, err := interactor.GetUsers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	return c.JSON(http.StatusOK, u)
}

func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	u, err := interactor.GetUser(c, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	return c.JSON(http.StatusOK, u)
}

func GetTasks(c echo.Context) error {
	t, err := interactor.GetTasks(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	return c.JSON(http.StatusOK, t)
}

func CreateTask(c echo.Context) error {
	task, err := interactor.BindCreateTask(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	if err = interactor.IsExistsUser(c, task.UserID); err != nil {
		return c.JSON(http.StatusNotFound, `{"message": "not found"}`)
	}
	if err = interactor.CreateTask(c, task.UserID, task.Title); err != nil {
		return c.JSON(http.StatusInternalServerError, `{"message": "internal server error"}`)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}
