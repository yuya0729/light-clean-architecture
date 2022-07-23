package gateway

import (
	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

// usecaseから呼ばれる
// SQL叩く
// entityに渡す
func GetUsers(c echo.Context) *entity.User {
	u := &entity.User{
		Id:   1,
		Name: "user1",
	}
	return u
}
