package gateway

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	myerror "github.com/yuya0729/light-clean-architecture/cmd/app/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"

	tasks "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway/tasks"
	users "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway/users"
)

var DB *sql.DB

// usecaseから呼ばれる
// SQL叩く
// entityに渡す

// users
func GetUsers(c echo.Context) ([]*entity.User, error) {
	return users.GetUsers(c, DB)
}

func GetUser(c echo.Context, userID int) (*entity.User, *myerror.MyError) {
	return users.GetUser(c, DB, userID)
}

// tasks
func GetTasks(c echo.Context) ([]*entity.Task, *myerror.MyError) {
	return tasks.GetTasks(c, DB)
}

func GetTask(c echo.Context, userID int, taskID int) (*entity.Task, *myerror.MyError) {
	return tasks.GetTask(c, DB, userID, taskID)
}

func CreateTask(c echo.Context, userID int, title string) *myerror.MyError {
	return tasks.CreateTask(c, DB, userID, title)
}

func UpdateTask(c echo.Context, userID int, title string, taskID int) *myerror.MyError {
	return tasks.UpdateTask(c, DB, userID, title, taskID)
}

func DeleteTask(c echo.Context, userID int, taskID int) *myerror.MyError {
	return tasks.DeleteTask(c, DB, userID, taskID)
}
