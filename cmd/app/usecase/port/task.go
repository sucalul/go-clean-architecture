package port

import (
	"context"

	"github.com/yuya0729/go-clean-architecture/tree/develop/cmd/app/entity"
)

type TaskInputPort interface {
	GetTasks(ctx context.Context)
}

type TaskOutputPort interface {
	Render(*entity.Task)
	RenderError(error)
}

// taskのCRUDに対するDB用ポート

type TaskRepository interface {
	GetTasks(ctx context.Context) (*entity.Task, error)
}
