package interactor

import (
	"context"

	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/port"
)

type Task struct {
	OutputPort port.TaskOutputPort
	TaskRepo   port.TaskRepository
}

// NewTaskInputPort はTaskInputPortを取得します．
func NewTaskInputPort(outputPort port.TaskOutputPort, taskRepository port.TaskRepository) port.TaskInputPort {
	return &Task{
		OutputPort: outputPort,
		TaskRepo:   taskRepository,
	}
}

// usecase.TaskInputPortを実装している
// GetTaskList は，TaskRepo.GetTaskListを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (t *Task) GetTaskList(ctx context.Context) {
	task, err := t.TaskRepo.GetTaskList(ctx)
	if err != nil {
		t.OutputPort.RenderError(err)
		return
	}
	t.OutputPort.Render(task)
}
