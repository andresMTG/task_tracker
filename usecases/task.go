package usecases

import (
	"fmt"
	"time"

	"github.com/andresMTG/task_tracker/repository"
)

// CreateTask creates a new task and add to our jsonFile
func CreateTask(description string, taskList []*repository.Task, fileName string) {

	newTask := repository.Task{
		Id:          autoIncrementId(taskList),
		Description: description,
		Status:      repository.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskList = append(taskList, &newTask)

	WriteFile(fileName, taskList)

	fmt.Printf("Task %v succesfully added to your list", newTask.Id)
}

// autoIncrementId is workaround to have auto incremented id in a non database env
func autoIncrementId(taskList []*repository.Task) int {
	lastId := 0
	for _, task := range taskList {
		if task.Id > lastId {
			lastId = task.Id
		}
	}
	return lastId + 1
}

// UpdateDescription change the description to a new one
func UpdateDescription(taskList []*repository.Task, id int, newDescription,fileName string){
	updatedTask, err := getTaskById(id,taskList)
	if err != nil {
		fmt.Println(err)
		return 
	}

	updatedTask.Description = newDescription
	updatedTask.UpdatedAt = time.Now()

	WriteFile(fileName, taskList)

	fmt.Printf("Task %v succesfully updated to your list", id)
}

// getTaskById is an auxiliar function to get expecified task
func getTaskById(id int, taskList []*repository.Task) (*repository.Task, error) {
	for _, task := range taskList {
		if task.Id == id {
			return task,nil;
		}
	}
	return &repository.Task{}, fmt.Errorf("Id dont exists error: ID - %v, dont exists", id)
}