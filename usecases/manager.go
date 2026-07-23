package usecases

import (
	"fmt"
	"os"
	"strconv"
)

// CommandManager is the main function that orchestrate the system
func CommandManager(command string, fileName string) {
	taskList := FileToListTask(fileName)

	if len(os.Args) <= 2 && command != "list" {
		command = "help"
	}
	
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
		DeleteTask(taskList, id, fileName)
	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(taskList, id, command, fileName)
	case "mark-in-done":
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(taskList, id, command, fileName)
	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		ShowTasks(taskList, filter)
	case "help":
		help()
	}
}

func help() {
	helpText :=
		`
task_tracker is a simple CLI tool for organizing your tasks.

USAGE:
	task_tracker <command> [argument]
COMMANDS:
	add <description>
		Create a new task.

	update <id> <description>
		Change the description of an existing task.

	delete <id>
		Remove an existing task.

	mark-in-progress <id>
		Mark a task as in progress.

	mark-in-done <id>
		Mark a task as done.

	list [filter]
		List all tasks. Optionally filter by: todo, in-progress, or done.

	help
		Show this help message.
`

	fmt.Println(helpText)
}
