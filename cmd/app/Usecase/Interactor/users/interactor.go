package users

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
// memo

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

func IsExistsUser(c echo.Context, userID int) error {
	if _, err := gateway.GetUser(c, userID); err != nil {
		return errors.New("not found")
	}
	return nil
}
