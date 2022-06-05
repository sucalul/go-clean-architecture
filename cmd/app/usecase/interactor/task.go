package interactor

import (
	"context"

	"github.com/yuya0729/echo-env/cmd/app/usecase/port"
)

type Task struct {
	OutputPort port.TaskOutputPort
	TaskRepo   port.TaskRepository
}

// NewTaskInputPortはTaskInputPortを取得
// ここちょっと違う
// func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
func NewTaskInputPort(outputPort port.TaskOutputPort, taskRepository port.TaskRepository) *port.TaskInputPort {
	return &Task{
		OutputPort: outputPort,
		TaskRepo:   taskRepository,
	}
}

func (t *Task) GetTasks(ctx context.Context) {
	task, err := t.TaskRepo.GetTasks(ctx)
	if err != nil {
		t.OutputPort.RenderError(err)
		return
	}
	t.OutputPort.Render(task)
}
