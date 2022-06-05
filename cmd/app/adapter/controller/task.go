package controller

import (
	"database/sql"
	"net/http"

	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/usecase/port"
)

type Task struct {
	OutputFactory func(w http.ResponseWriter) port.TaskOutputPort
	// -> presenter.NewTaskOutputPort
	InputFactory func(o port.TaskOutputPort, u port.TaksRepository) port.TaskInputPort
	// -> interactor.NewTaksInputPort
	RepoFactory func(c *sql.DB) port.TaskRepository
	// -> gateway.NewTaksRepository
	Conn *sql.DB
}

func (t *Task) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	println(ctx)
	outputPort := t.OutputFactory(w)
	repository := t.RepoFactory(t.Conn)
	inputPort := t.InputFactory(outputPort, repository)
	inputPort.GetTasks(ctx)
}
