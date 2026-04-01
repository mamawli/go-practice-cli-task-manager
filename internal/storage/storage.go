package storage

import (
	"fmt"
	"os"
	"path/filepath"
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

func WriteToDb(content []byte) {

	path := getDbPath()
	_ = os.Remove(path)
	err := os.WriteFile(path, content, os.FileMode(777))
	if err != nil {
		fmt.Println("Failed to write file", err)
		os.Exit(1)
	}
}

func ReadTasks() []byte {
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
