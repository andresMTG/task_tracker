package usecases

import (
	"fmt"
	"os"
	"strconv"
)

// CommandManager is the main function that orchestrate the system
func CommandManager(command string, fileName string) {
	taskList := FileToListTask(fileName)
	switch command {
	case "add":
		description := os.Args[2]
		CreateTask(description, taskList, fileName)
	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		UpdateDescription(taskList, id, description, fileName)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		DeleteTask(taskList , id , fileName )
	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(taskList , id , command, fileName )
	case "mark-in-done":
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(taskList , id , command, fileName )
	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		ShowTasks(taskList, filter)
	case "help":
		fmt.Println("you need to choose what to do")
	}
}
