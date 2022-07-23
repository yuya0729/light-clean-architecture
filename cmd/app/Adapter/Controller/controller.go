package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// driverで定義されたエンドポイントの関数を定義する
func GetUsers(c echo.Context) error {
	u, err := interactor.GetUsers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `{"message": "bad request"}`)
	}
	return c.JSON(http.StatusOK, u)
}
