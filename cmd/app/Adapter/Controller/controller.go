package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

// driverで定義されたエンドポイントの関数を定義する
func GetUsers(c echo.Context) error {
	u := &entity.User{
		Id:   2,
		Name: "user1",
	}
	return c.JSON(http.StatusOK, u)
}
