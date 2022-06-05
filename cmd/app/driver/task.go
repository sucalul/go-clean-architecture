package drivers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/adapter/controller"
	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/adapter/geteway"
	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/adapter/presenter"
	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/usecase/interactor"
)

func Serve() {
	dsn := fmt.Sprint("user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	task := controller.Task{
		OutputFactory: presenter.NewTaskOutputPort,
		InputFactory:  interactor.NewTaskInputPort,
		RepoFactory:   geteway.NewTaskRepository,
		Conn:          conn,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/tasks", task.GetTasks)
	e.Start(":8080")
}
