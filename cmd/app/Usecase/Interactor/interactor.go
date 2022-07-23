package interactor

import (
	"github.com/labstack/echo/v4"
	gateway "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い
func GetUsers(c echo.Context) ([]*entity.User, error) {
	u, err := gateway.GetUsers(c)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(c echo.Context, userID int) (*entity.User, error) {
	u, err := gateway.GetUser(c, userID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetTasks(c echo.Context) ([]*entity.Task, error) {
	t, err := gateway.GetTasks(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}
