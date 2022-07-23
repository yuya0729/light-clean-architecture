package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

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
