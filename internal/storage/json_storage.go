package storage

import (
	"encoding/json"
	"fmt"
	"go-practice-cli-task-manager/internal/task"
	"os"
	"path/filepath"
	"sync"
)

type JsonStorage struct {
	mu sync.Mutex
}

func NewJsonStorage() JsonStorage {
	jsonPath := getDbPath()
	_, err := os.Stat(jsonPath)
	if os.IsNotExist(err) {
		_, err := os.Create(jsonPath)
		if err != nil {
			fmt.Println("Failed to create json file", err)
			os.Exit(1)
		}
	}
	return JsonStorage{}
}

func (j *JsonStorage) List() ([]task.Task, error) {
	content := readDB()
	if len(content) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (j *JsonStorage) Add(t *task.Task) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	tasks, err := j.List()
	if err != nil {
		fmt.Println("Failed to add task")
		return err
	}

	t.Id = getNextId(tasks) + 1
	tasks = append(tasks, *t)
	err = writeToDb(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonStorage) GetByID(id int64) (task.Task, error) {
	tasks, err := j.List()
	if err != nil {
		return task.Task{}, err
	}
	for _, t := range tasks {
		if t.Id == id {
			return t, nil
		}
	}
	return task.Task{}, fmt.Errorf("task with id %d not found", id)
}

func (j *JsonStorage) MarkDone(id int64) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	tasks, err := j.List()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].CompleteTask()
			break
		}
	}
	err = writeToDb(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonStorage) Delete(id int64) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	tasks, err := j.List()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	err = writeToDb(tasks)
	if err != nil {
		return err
	}
	return nil
}

func getNextId(tasks []task.Task) int64 {
	var nextId int64 = 0
	for _, ts := range tasks {
		if ts.Id > nextId {
			nextId = ts.Id
		}
	}
	return nextId
}

func readDB() []byte {
	content := openDb()
	if len(content) == 0 {
		return nil
	}

	return content
}

func openDb() []byte {
	path := getDbPath()
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read db content", err)
	}

	return content
}

func getDbPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	root := filepath.Dir(filepath.Dir(cwd))
	filePath := filepath.Join(root, "tasks.json")
	return filePath
}

func writeToDb(t []task.Task) error {

	content, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to marshal json")
		return err
	}
	path := getDbPath()
	_ = os.Remove(path)
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Println("Failed to write file")
		return err
	}
	return nil
}
