package driver

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/yuya0729/go-clean-architecture/cmd/app/adapter/controller"
	"github.com/yuya0729/go-clean-architecture/cmd/app/adapter/gateway"
	"github.com/yuya0729/go-clean-architecture/cmd/app/adapter/presenter"
	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/interactor"
)

func Serve() {
	dsn := "user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	task := controller.Task{
		OutputFactory: presenter.NewTaskOutputPort,
		InputFactory:  interactor.NewTaskInputPort,
		RepoFactory:   gateway.NewTaskRepository,
		Conn:          conn,
	}
	// task一覧の取得
	http.HandleFunc("/api/v1/tasks", task.GetTaskList)
	// TODO: high
	// task作成のエンドポイント
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
