package usecases

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/andresMTG/task_tracker/repository"
)


func CreateTask(description string, taskList []repository.Task, fileName string) {

	newTask := repository.Task{
		Id:          autoIncrementId(taskList),
		Description: description,
		Status:      repository.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskList = append(taskList, newTask)

	jsonToSave, err2 := json.Marshal(taskList)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := os.WriteFile(fileName, jsonToSave, 7777)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Printf("Task %v succesfully added to your list", newTask.Id)
}


func autoIncrementId(taskList []repository.Task) int {
	lastId := 0
	for _, task := range taskList {
		if task.Id > lastId {
			lastId = task.Id
		}
	}
	return lastId + 1
}