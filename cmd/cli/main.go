package main

import (
	"flag"
	"fmt"
	"go-practice-cli-task-manager/internal/storage"
	"go-practice-cli-task-manager/internal/task"
)

func main() {

	addFlag := flag.String("add", "", "gimme the title of a task")
	listFlag := flag.Bool("list", false, "List of tasks")
	getTaskFlag := flag.Int64("get", 0, "Get task detail by id")
	doneTaskFlag := flag.Int64("done", 0, "Done task by id")
	deleteTaskFlag := flag.Int64("delete", 0, "delete task by id")

	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Invalid Flag")
		flag.Usage()
		return
	}
	jsonStorage := storage.NewJsonStorage()

	if *listFlag {
		tasks, err := jsonStorage.List()
		if err != nil {
			fmt.Println("Failed to get tasks", err)
			return
		}
		fmt.Println(tasks)
		return
	}

	if *addFlag != "" {
		newTask := task.NewTask(*addFlag)
		err := jsonStorage.Add(newTask)
		if err != nil {
			fmt.Println("Failed to add new task", err)
		}
		return
	}

	if *getTaskFlag != 0 {
		t, err := jsonStorage.GetByID(*getTaskFlag)
		if err != nil {
			fmt.Println("failed to get task by id", *getTaskFlag, err)
			return
		}
		fmt.Println(t)
		return
	}

	if *doneTaskFlag != 0 {
		err := jsonStorage.MarkDone(*doneTaskFlag)
		if err != nil {
			fmt.Println("Failed to mark task to done", err)
		}
		return
	}

	if *deleteTaskFlag != 0 {
		err := jsonStorage.Delete(*deleteTaskFlag)
		if err != nil {
			fmt.Println("Failed to delete task", err)
		}
		return
	}
}
