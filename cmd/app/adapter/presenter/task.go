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

// NewTaskOutputPort はTaskOutputPortを取得します．
func NewTaskOutputPort(w http.ResponseWriter) port.TaskOutputPort {
	return &Task{
		w: w,
	}
}

// usecase.TaskOutputPortを実装している
// Render はNameを出力します．
func (t *Task) Render(task *entity.Task) {
	t.w.WriteHeader(http.StatusOK)
	// httpでentity.Task.Nameを出力
	fmt.Fprint(t.w, task.Name)
}

// RenderError はErrorを出力します．
func (t *Task) RenderError(err error) {
	t.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(t.w, err)
}
