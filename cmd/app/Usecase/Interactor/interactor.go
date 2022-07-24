package interactor

import (
	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"

	tasks "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor/tasks"
	users "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor/users"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetUsers(c echo.Context) ([]*entity.User, error) {
	return users.GetUsers(c)
}

func GetUser(c echo.Context, userID int) (*entity.User, error) {
	return users.GetUser(c, userID)
}

func IsExistsUser(c echo.Context, userID int) error {
	return users.IsExistsUser(c, userID)
}

func GetTasks(c echo.Context) ([]*entity.Task, error) {
	return tasks.GetTasks(c)
}

func BindCreateUpdateTask(c echo.Context) (*entity.CreateTask, error) {
	return tasks.BindCreateUpdateTask(c)
}

func CreateTask(c echo.Context, userID int, title string) error {
	return tasks.CreateTask(c, userID, title)
}

func UpdateTask(c echo.Context, userID int, title string, taskID int) error {
	return tasks.UpdateTask(c, userID, title, taskID)
}

func IsExistsTask(c echo.Context, taskID int) error {
	return tasks.IsExistsTask(c, taskID)
}
