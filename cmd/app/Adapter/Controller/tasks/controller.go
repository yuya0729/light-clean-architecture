package tasks

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// TODO:
// err != nil{return c.JSON(statuscode, err)}
// みたいな感じにする？

func GetTasks(c echo.Context) error {
	t, err := interactor.GetTasks(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	return c.JSON(http.StatusOK, t)
}

func CreateTask(c echo.Context) error {
	task, err := interactor.BindCreateUpdateTask(c)
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

func UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	task, err := interactor.BindCreateUpdateTask(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	if err = interactor.IsExistsTask(c, task.UserID, taskID); err != nil {
		return c.JSON(http.StatusNotFound, `{"message": "not found"}`)
	}
	if err = interactor.IsExistsUser(c, task.UserID); err != nil {
		return c.JSON(http.StatusNotFound, `{"message": "not found"}`)
	}

	if err = interactor.UpdateTask(c, task.UserID, task.Title, taskID); err != nil {
		return c.JSON(http.StatusInternalServerError, `{"message": "internal server error"}`)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}

func DeleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	if err = interactor.IsExistsTask(c, userID, taskID); err != nil {
		return c.JSON(http.StatusNotFound, `{"message": "not found"}`)
	}
	if err = interactor.IsExistsUser(c, userID); err != nil {
		return c.JSON(http.StatusNotFound, `{"message": "not found asdf"}`)
	}

	if err = interactor.DeleteTask(c, userID, taskID); err != nil {
		return c.JSON(http.StatusInternalServerError, `{"message": "internal server error"}`)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}
