package tasks

import (
	"errors"

	"github.com/labstack/echo/v4"
	gateway "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetTasks(c echo.Context) ([]*entity.Task, error) {
	t, err := gateway.GetTasks(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func BindCreateTask(c echo.Context) (*entity.CreateTask, error) {
	task := entity.CreateTask{}
	if err := c.Bind(&task); err != nil {
		return nil, errors.New("bad request")
	}
	return &task, nil
}

func IsExistsUser(c echo.Context, userID int) error {
	if _, err := gateway.GetUser(c, userID); err != nil {
		return errors.New("not found")
	}
	return nil
}

func CreateTask(c echo.Context, userID int, title string) error {
	if err := gateway.CreateTask(c, userID, title); err != nil {
		return err
	}
	return nil
}
