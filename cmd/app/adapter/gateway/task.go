package gateway

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/yuya0729/go-clean-architecture/cmd/app/entity"
	"github.com/yuya0729/go-clean-architecture/cmd/app/usecase/port"
)

type TaskRepository struct {
	conn *sql.DB
}

// NewTaskRepository はTaskRepositoryを返します．
func NewTaskRepository(conn *sql.DB) port.TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}

// GetDBConn はconnectionを取得します．
func (t *TaskRepository) GetDBConn() *sql.DB {
	return t.conn
}

// GetTaskList はDBからデータを取得します．
func (t *TaskRepository) GetTaskList(ctx context.Context) (*entity.Task, error) {
	// TODO: high
	// jsonで返す
	conn := t.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM tasks")
	task := entity.Task{}
	err := row.Scan(&task.ID, &task.Name)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal Server Error. adapter/gateway/GetTaskList")
	}
	return &task, nil
}
