package usecases

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/andresMTG/task_tracker/repository"
)

// FileExists is simple validation function to validate if we have
// especified file in the directory
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if errors.Is(err, fs.ErrNotExist){
		return false
	}

	return true
}

// CreateFile creates a file if notExists, if exists they truncate
func CreateFile(fileName string) (*os.File) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return file
	}
	defer file.Close()
	return file
}

// WriteFile is a core function that write info into our json file
func WriteFile(fileName string, taskList []*repository.Task) {
	
	jsonToSave, marshalError := json.Marshal(taskList)
	if marshalError != nil {
		log.Println(marshalError)
	}

	writeError := os.WriteFile(fileName, jsonToSave, 7777)
	if writeError != nil {
		log.Println(writeError)
	}
}

// readFile is an auxiliar function to take the data on our jsonfile
func readFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	
	return file
}

// FileToListTask has the function to converte the json data to a slice of tasks
func FileToListTask(fileName string) []*repository.Task {
	var listTasks []*repository.Task

	file := readFile(fileName)

	if len(file) == 0 {
		return listTasks
	}

	err := json.Unmarshal(file, &listTasks)
	if err != nil {
		log.Println(err)
	}
	return listTasks
}
