package presenter

import (
	"fmt"
	"net/http"

	"github.com/yuya0729/go-clean-architecture/cmd/app/entity"
	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/port"
)

type Task struct {
	w http.ResponseWriter
}

func NewTaskOutputPort(w http.ResponseWriter) port.TaskOutputPort {
	return &Task{
		w: w,
	}
}

func (t *Task) Render(task *entity.Task) {
	t.w.WriteHeader(http.StatusOK)
	fmt.Fprint(t.w, task.Task)
}
