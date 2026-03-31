package main

import (
	"fmt"
	jsonstorage "json_storage/internal/storage"
	"json_storage/internal/task"
	"os"
	"strings"
)

func main() {

	jsonstorage.AutoCreateDB()
	action := os.Args[1]
	strings.ToLower(action)

	switch action {
	case "add":
		v := os.Args[2]
		if v == "" {
			fmt.Println("title must be filled!")
		}
		newTask := task.NewTask(v)
		jsonstorage.AddTask(newTask)
	case "list":
		fmt.Println(jsonstorage.ReadTasks())

	case "get":
		v := os.Args[2]
		if v == "" {
			fmt.Println("id must filled")
		}
		fmt.Println(jsonstorage.GetTaskById(v))
	case "done":
		v := os.Args[2]
		if v == "" {
			fmt.Println("id must filled")
		}
		jsonstorage.DoneTask(v)

	case "delete":
		v := os.Args[2]
		if v == "" {
			fmt.Println("id must filled")
		}
		jsonstorage.DeleteTask(v)
	}
}
