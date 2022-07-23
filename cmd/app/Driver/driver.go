package driver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Controller"
)

// routerとかinit dbとか
func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// APIルーティング
	e.GET("/api/users", controller.GetUsers)

	e.Start(":8080")
}
