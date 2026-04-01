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

	used := false
	flag.VisitAll(func(f *flag.Flag) {
		used = true
	})

	if !used {
		fmt.Println("Invalid Flag")
		flag.Usage()
		return
	}

	storage.AutoCreateDB()

	flag.Parse()

	if *listFlag {
		fmt.Println(task.ReadTasks())
		return
	}

	if *addFlag != "" {
		newTask := task.NewTask(*addFlag)
		task.AddTask(newTask)
		return
	}

	if *getTaskFlag != 0 {
		fmt.Println(task.GetTaskById(*getTaskFlag))
		return
	}

	if *doneTaskFlag != 0 {
		task.CompleteTask(*doneTaskFlag)
		return
	}

	if *deleteTaskFlag != 0 {
		task.DeleteTask(*deleteTaskFlag)
		return
	}
}
