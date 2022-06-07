package port

import (
	"context"

	"github.com/yuya0729/go-clean-architecture/cmd/app/entity"
)

type TaskInputPort interface {
	GetTaskList(ctx context.Context)
}

type TaskOutputPort interface {
	Render(*entity.Task)
	RenderError(error)
}

// taskのCRUDに対するDB用のポート
type TaskRepository interface {
	GetTaskList(ctx context.Context) (*entity.Task, error)
}
