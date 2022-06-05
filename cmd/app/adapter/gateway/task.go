package geteway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/yuya0729/go-clean-architecture/cmd/app/entity"
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

func (t *TaskRepository) GetDBConn() *sql.DB {
	return t.conn
}

func (t *TaskRepository) GetTasks(ctx context.Context) (*entity.Task, error) {
	conn := t.GetDBConn()
	row, err := conn.Query("SELECT * FROM Tasks")
	task := entity.Task{}

	err := row.Scan(&task.ID, &task.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task Not Found.")
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetTasks")
	}
	return &task, nil
}
