package task

import (
	"time"
)

type Task struct {
	Id        int64      `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	Created   time.Time  `json:"created"`
	Completed *time.Time `json:"completed"`
}

func NewTask(title string) *Task {

	return &Task{
		Title:   title,
		Done:    false,
		Created: time.Now(),
	}
}

func (t *Task) CompleteTask() {
	now := time.Now()
	t.Completed = &now
	t.Done = true
}
