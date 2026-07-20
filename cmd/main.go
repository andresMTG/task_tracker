package main

import (
	"os"

	"github.com/andresMTG/task_tracker/usecases"
)

const listTaskName = "listTask.json"

func main() {
	existFile := usecases.FileExists(listTaskName)
	
	if existFile != true {
		usecases.CreateFile(listTaskName)
	}
	
	command := "help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	
	usecases.CommandManager(command, listTaskName)

}
