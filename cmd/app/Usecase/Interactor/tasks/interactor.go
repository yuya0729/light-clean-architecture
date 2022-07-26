package tasks

import (
	"github.com/labstack/echo/v4"
	gateway "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway"
	myerror "github.com/yuya0729/light-clean-architecture/cmd/app/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetTasks(c echo.Context) ([]*entity.Task, *myerror.MyError) {
	t, err := gateway.GetTasks(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func BindCreateUpdateTask(c echo.Context) (*entity.CreateTask, *myerror.MyError) {
	task := entity.CreateTask{}
	if err := c.Bind(&task); err != nil {
		return nil, myerror.New(400, err.Error())
	}
	return &task, nil
}

func IsExistsTask(c echo.Context, userID int, taskID int) *myerror.MyError {
	if _, err := gateway.GetTask(c, userID, taskID); err != nil {
		return err
	}
	return nil
}

func CreateTask(c echo.Context, userID int, title string) *myerror.MyError {
	if err := gateway.CreateTask(c, userID, title); err != nil {
		return err
	}
	return nil
}

func UpdateTask(c echo.Context, userID int, title string, taskID int) *myerror.MyError {
	if err := gateway.UpdateTask(c, userID, title, taskID); err != nil {
		return err
	}
	return nil
}

func DeleteTask(c echo.Context, userID int, taskID int) *myerror.MyError {
	if err := gateway.DeleteTask(c, userID, taskID); err != nil {
		return err
	}
	return nil
}
