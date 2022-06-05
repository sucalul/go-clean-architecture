package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

func main() {
	// Echoのインスタンス
	e := echo.New()

	// ミドルウェア類
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/api/v1/tasks", GetTasks())

	// サーバー起動
	e.Start(":8080")
}

type Task struct {
	Id   int
	Task string
}

var Db *sql.DB

func init() {
	var err error
	// これめんどい
	// https: //qiita.com/mabubu0203/items/5cdff1caf2b024df1d95
	Db, err = sql.Open("postgres", "user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		task := Task{}
		tasks := []*Task{}

		rows, err := Db.Query("select * from tasks")
		if err != nil {
			return errors.Wrapf(err, "connot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&task.Id, &task.Task); err != nil {
				return errors.Wrapf(err, "connot connect SQL")
			}
			tasks = append(tasks, &Task{Id: task.Id, Task: task.Task})
		}

		return c.JSON(http.StatusOK, tasks)
	}
}
