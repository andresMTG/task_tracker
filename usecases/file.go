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

func readFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	
	return file
}

func FileToListTask(fileName string) []repository.Task {
	var listTasks []repository.Task

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