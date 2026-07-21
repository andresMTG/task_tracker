package usecases

import (
	"fmt"
	"slices"
	"time"

	"github.com/andresMTG/task_tracker/repository"
)

const (
	ColorReset = "\033[0m"
	ColorGreen = "\033[32m"
)

var status = []string{repository.DONE, repository.IN_PROGRESS, repository.TODO}

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
func UpdateDescription(taskList []*repository.Task, id int, newDescription, fileName string) {
	updatedTask, _, err := getTaskById(id, taskList)
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedTask.Description = newDescription
	updatedTask.UpdatedAt = time.Now()

	WriteFile(fileName, taskList)

	fmt.Printf("Task %v succesfully updated to your list", id)
}

// getTaskById is an auxiliar function to get expecified task and their indice
func getTaskById(id int, taskList []*repository.Task) (*repository.Task, int, error) {
	var returnIndice int
	for indice, task := range taskList {
		if task.Id == id {
			returnIndice = indice
			return task, returnIndice, nil
		}
	}
	return &repository.Task{}, returnIndice, fmt.Errorf("Id dont exists error: ID - %v, dont exists", id)
}

// DeleteTask remove expecific task from our jason file by id
func DeleteTask(taskList []*repository.Task, id int, fileName string) {
	_, indice, err := getTaskById(id, taskList)
	if err != nil {
		fmt.Println(err)
		return
	}

	taskList = append(taskList[:indice], taskList[indice+1:]...)

	WriteFile(fileName, taskList)

	fmt.Printf("Task %v succesfully deleted from your list", id)
}

// UpdateStatus change the status of the expecific task
func UpdateStatus(taskList []*repository.Task, id int, status, fileName string) {
	updatedTask, _, err := getTaskById(id, taskList)
	if err != nil {
		fmt.Println(err)
		return
	}

	newStatus := repository.DONE

	if status == "mark-in-progress" {
		newStatus = repository.IN_PROGRESS
	}

	updatedTask.Status = newStatus
	updatedTask.UpdatedAt = time.Now()

	WriteFile(fileName, taskList)

	fmt.Printf("Task %v current status is: %v", id, newStatus)
}

// ShowTasks use a filter to decide each type of task have to show
func ShowTasks(taskList []*repository.Task, filter string) {

	filteredTasks := []*repository.Task{}

	if slices.Contains(status, filter) {
		formatTaskList(filterTasks(filter, filteredTasks, taskList))
		return
	}

	if filter != "" {
		fmt.Printf("Choosen parameter '%v' is not a valid parameter. Valid parameters: %v\t%v\t%v", filter, status[0], status[1], status[2])
		return
	}

	formatTaskList(taskList)
}

// filterTasks recieve all tasks and filter then by status
func filterTasks(filter string, taskfiltered, taskList []*repository.Task) []*repository.Task {
	for _, task := range taskList {
		if task.Status == filter {
			taskfiltered = append(taskfiltered, task)
		}
	}
	return taskfiltered
}

// formatTaskList auxiliar function to define how the task data you showup
func formatTaskList(taskList []*repository.Task) {
	for _, task := range taskList {
		fmt.Printf("ID: %s%v%s \tStatus: %s%v%s \n\tDescription: %s%v%s\n",
			ColorGreen, task.Id, ColorReset, ColorGreen, task.Status, ColorReset, ColorGreen, task.Description, ColorReset)
	}
}
