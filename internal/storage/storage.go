package storage

import "go-practice-cli-task-manager/internal/task"

type Store interface {
	Add(t *task.Task) error
	List() ([]task.Task, error)
	GetByID(id int64) (task.Task, error)
	MarkDone(id int64) error
	Delete(id int64) error
}
