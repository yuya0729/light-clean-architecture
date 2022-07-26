package tasks

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	myerror "github.com/yuya0729/light-clean-architecture/cmd/app/Driver/error"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// TODO:
// interface導入？

// TODO:
// controllerでやっているバリデーションを各層に分散
// ref: https://qiita.com/nakaakist/items/11195ac5290450fbc9f9

// TODO:
// errorの体系化

func GetTasks(c echo.Context) error {
	t, err := interactor.GetTasks(c)
	if err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusOK, t)
}

func CreateTask(c echo.Context) error {
	task, err := interactor.BindCreateUpdateTask(c)
	if err != nil {
		switch err.Code {
		case 400:
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	if err := interactor.IsExistsUser(c, task.UserID); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	if err = interactor.CreateTask(c, task.UserID, task.Title); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusOK, `{"code": 200, "message": "ok"}`)
}

func UpdateTask(c echo.Context) error {
	taskID, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		msg := myerror.New(400, e.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}
	task, err := interactor.BindCreateUpdateTask(c)
	if err != nil {
		switch err.Code {
		case 400:
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	if err = interactor.IsExistsTask(c, task.UserID, taskID); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	if err = interactor.UpdateTask(c, task.UserID, task.Title, taskID); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusOK, `{"code": 200, "message": "ok"}`)
}

func DeleteTask(c echo.Context) error {
	taskID, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		msg := myerror.New(400, e.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}
	userID, e := strconv.Atoi(c.QueryParam("user_id"))
	if e != nil {
		msg := myerror.New(400, e.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}
	if err := interactor.IsExistsTask(c, userID, taskID); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	if err := interactor.DeleteTask(c, userID, taskID); err != nil {
		switch err.Code {
		case 404:
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}
