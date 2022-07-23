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
func GetUsers(c echo.Context) *entity.User {
	return gateway.GetUsers(c)
}
