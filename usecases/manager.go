package usecases

import (
	"fmt"
	"os"
)

func CommandManager(command string, fileName string) {
	switch command {
	case "add":
		description := os.Args[2]
		listTask := FileToListTask(fileName)
		CreateTask(description,listTask,fileName)
	case "update":
		fmt.Println("Task updated successfully")
	case "delete":
		fmt.Println("task deleted successfully")
	case "mark-in-progress":
		fmt.Println("mark-in-progress")
	case "mark-in-done":
		fmt.Println("mark-in-done")
	case "list":
		fmt.Println("list")
	case "help":
		fmt.Println("you need to choose what to do")
	}
}