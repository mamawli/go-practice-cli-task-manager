package task

import (
	"encoding/json"
	"fmt"
	"go-practice-cli-task-manager/internal/storage"
	"math"
	"os"
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

func ReadTasks() []Task {
	return readTasks()
}

func GetTaskById(id int64) Task {

	tasks := readTasks()
	for _, t := range tasks {
		if t.Id == id {
			return t
		}
	}
	return Task{}
}

func CompleteTask(id int64) {
	tasks := readTasks()
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].completeTask()
			break
		}
	}
	writeToDb(tasks)
}

func DeleteTask(id int64) {
	tasks := readTasks()
	for i := range tasks {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	writeToDb(tasks)
}

func AddTask(t *Task) {

	tasks := readTasks()
	var maxId int64 = 0
	for _, t := range tasks {
		maxId = int64(math.Max(float64(t.Id), float64(maxId)))
	}
	t.Id = maxId + 1
	tasks = append(tasks, *t)
	writeToDb(tasks)
}

func (t *Task) completeTask() {
	now := time.Now()
	t.Completed = &now
	t.Done = true
}

func writeToDb(t []Task) {

	content, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to marshal json", err)
	}
	storage.WriteToDb(content)
}

func readTasks() []Task {
	content := storage.ReadTasks()
	if len(content) == 0 {
		return nil
	}
	var tasks []Task
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("Failed to map data", err)
		os.Exit(1)
	}
	return tasks
}
