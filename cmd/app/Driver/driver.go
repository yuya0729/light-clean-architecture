package driver

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Controller"
	gateway "github.com/yuya0729/light-clean-architecture/cmd/app/Adapter/Gateway"
)

// routerとかdb接続とか
func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// db接続
	var err error
	dsn := "user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable"
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// gateway DBにDBを渡す
	gateway.DB = DB

	// APIルーティング
	api := e.Group("/api")
	api.GET("/users", controller.GetUsers)
	api.GET("/users/:id", controller.GetUser)

	e.Start(":8080")
}
