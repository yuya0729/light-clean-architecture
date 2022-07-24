package tasks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// TODO:
// err != nil{return c.JSON(statuscode, err)}
// みたいな感じにする？

// TODO:
// interface導入？

// TODO:
// controllerでやっているバリデーションを各層に分散

func GetTasks(c echo.Context) error {
	t, err := interactor.GetTasks(c)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, t)
}

func CreateTask(c echo.Context) error {
	task, err := interactor.BindCreateUpdateTask(c)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	if err = interactor.IsExistsUser(c, task.UserID); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusNotFound, msg)
	}
	if err = interactor.CreateTask(c, task.UserID, task.Title); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}

func UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	task, err := interactor.BindCreateUpdateTask(c)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	if err = interactor.IsExistsTask(c, task.UserID, taskID); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusNotFound, msg)
	}
	if err = interactor.UpdateTask(c, task.UserID, task.Title, taskID); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}

func DeleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	if err = interactor.IsExistsTask(c, userID, taskID); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusNotFound, msg)
	}

	if err = interactor.DeleteTask(c, userID, taskID); err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}
