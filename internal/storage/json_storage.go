package json_storage

import (
	"encoding/json"
	"fmt"
	"json_storage/internal/task"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func AutoCreateDB() {

	jsonPath := getDbPath()
	_, err := os.Stat(jsonPath)
	if os.IsNotExist(err) {
		_, err := os.Create(jsonPath)
		if err != nil {
			fmt.Println("Failed to create json file", err)
			os.Exit(1)
		}
	}
}

func getDbPath() string {
	dir, err := findProjectRoot()
	if err != nil {
		fmt.Printf("Failed to get current working root directory: %v\n\n", err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s/%s", dir, "tasks.json")

}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}

func GetTaskById(strId string) *task.Task {

	tasks := ReadTasks()
	for _, t := range tasks {
		if t.Id == convertIntId(strId) {
			return &t
		}
	}
	return nil
}

func DoneTask(strId string) {
	tasks := ReadTasks()
	for i, t := range tasks {
		if t.Id == convertIntId(strId) {
			tasks[i].Complete()
			break
		}
	}
	writeToDb(tasks)
}

func DeleteTask(strId string) {
	id := convertIntId(strId)
	tasks := ReadTasks()
	for i := range tasks {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	writeToDb(tasks)
}

func AddTask(t *task.Task) {

	tasks := ReadTasks()
	if len(tasks) != 0 {
		taskCopy := make([]task.Task, len(tasks))
		copy(taskCopy, tasks)
		sort.Slice(taskCopy, func(i, j int) bool {
			return taskCopy[i].Id > taskCopy[j].Id
		})

		t.Id = taskCopy[0].Id + 1
	} else if len(tasks) == 0 {
		t.Id = 1
	}

	tasks = append(tasks, *t)
	writeToDb(tasks)
}

func writeToDb(tasks []task.Task) {

	path := getDbPath()

	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Failed to marshal json", err)
	}
	_ = os.Remove(path)
	err = os.WriteFile(path, content, os.FileMode(777))
	if err != nil {
		fmt.Println("Failed to write file", err)
		os.Exit(1)
	}
}

func ReadTasks() []task.Task {
	content := openDb()
	if len(content) == 0 {
		return nil
	}
	var tasks []task.Task
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("Invalid Data", err)
		os.Exit(1)
	}

	return tasks
}

func openDb() []byte {
	path := getDbPath()
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read db content", err)
	}

	return content
}

func convertIntId(strId string) int64 {

	id, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println(" Id must be integer")
		os.Exit(1)
	}

	return int64(id)
}
