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
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, errors.New("connot connect SQL")
		}
		users = append(users, &entity.User{Id: user.Id, Name: user.Name})
	}
	return users, nil
}

func GetUser(c echo.Context, userID int) (*entity.User, error) {
	user := &entity.User{}
	err := DB.QueryRow("SELECT id, name FROM users WHERE id = $1", userID).Scan(&user.Id, &user.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func GetTasks(c echo.Context) ([]*entity.Task, error) {
	task := entity.Task{}
	tasks := []*entity.Task{}
	rows, err := DB.Query(`
		SELECT
			tasks.id as id,
			tasks.title,
			users.name
		FROM
			tasks
			JOIN
				users
			ON users.id = tasks.user_id
	`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.Name); err != nil {
			return nil, errors.New("connot connect SQL")
		}
		tasks = append(tasks, &entity.Task{Id: task.Id, Title: task.Title, Name: task.Name})
	}
	return tasks, nil
}

// TODO: log形式の統一

func CreateTask(c echo.Context, userID int, title string) error {
	ins, err := DB.Prepare("INSERT INTO tasks(user_id, title) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return err
	}
	ins.Exec(userID, title)
	return nil
}
