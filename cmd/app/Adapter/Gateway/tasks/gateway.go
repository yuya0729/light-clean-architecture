package tasks

import (
	"database/sql"
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"
)

func GetTasks(c echo.Context, DB *sql.DB) ([]*entity.Task, error) {
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
		if err := rows.Scan(&task.ID, &task.Title, &task.Name); err != nil {
			return nil, errors.New("connot connect SQL")
		}
		tasks = append(tasks, &entity.Task{ID: task.ID, Title: task.Title, Name: task.Name})
	}
	return tasks, nil
}

func CreateTask(c echo.Context, DB *sql.DB, userID int, title string) error {
	ins, err := DB.Prepare("INSERT INTO tasks(user_id, title) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return err
	}
	ins.Exec(userID, title)
	return nil
}
