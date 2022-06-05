package geteway

import (
	"database/sql"

	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/port"
)

type TaskRepository struct {
	conn *sql.DB
}

func NewTaskRepository(conn *sql.DB) *port.TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}
