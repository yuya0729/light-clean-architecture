package gateway

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

var DB *sql.DB

// usecaseから呼ばれる
// SQL叩く
// entityに渡す
func GetUsers(c echo.Context) ([]*entity.User, error) {
	user := entity.User{}
	users := []*entity.User{}
	rows, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal Server Error. adapter/gateway/GetTaskList")
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, errors.New("connot connect SQL")
		}
		users = append(users, &entity.User{Id: user.Id, Name: user.Name})
	}
	return users, nil
}
