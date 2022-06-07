package controller

import (
	"database/sql"
	"net/http"

	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/port"
)

type Task struct {
	OutputFactory func(w http.ResponseWriter) port.TaskOutputPort
	// -> presenter.NewTaskOutputPort
	InputFactory func(o port.TaskOutputPort, u port.TaskRepository) port.TaskInputPort
	// -> interactor.NewTaskInputPort
	RepoFactory func(c *sql.DB) port.TaskRepository
	// -> gateway.NewTaskRepository
	Conn *sql.DB
}

func (t *Task) GetTaskList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := t.OutputFactory(w)
	repository := t.RepoFactory(t.Conn)
	inputPort := t.InputFactory(outputPort, repository)
	inputPort.GetTaskList(ctx)
}
