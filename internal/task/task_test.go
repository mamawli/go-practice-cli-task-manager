package task_test

import (
	"go-practice-cli-task-manager/internal/task"
	"testing"
)

func TestNewTask(t *testing.T) {
	newTask := task.NewTask("Test")

	if newTask.Done {
		t.Error("Task cannot be done when created")
	}

	if newTask.Completed != nil {
		t.Error("Task Completed Date must be nil when created")
	}
}
