package controller

import (
	"github.com/labstack/echo/v4"
	tasks "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Controller/tasks"
	users "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Controller/users"
)

// driverで定義されたエンドポイントの関数を定義する

// users
func GetUsers(c echo.Context) error {
	return users.GetUsers(c)
}

func GetUser(c echo.Context) error {
	return users.GetUser(c)
}

//
func GetTasks(c echo.Context) error {
	return tasks.GetTasks(c)
}

func CreateTask(c echo.Context) error {
	return tasks.CreateTask(c)
}
