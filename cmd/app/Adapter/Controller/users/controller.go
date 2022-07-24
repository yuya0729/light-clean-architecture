package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor"
)

// driverで定義されたエンドポイントの関数を定義する
func GetUsers(c echo.Context) error {
	u, err := interactor.GetUsers(c)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, u)
}

func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	u, err := interactor.GetUser(c, userID)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, u)
}
