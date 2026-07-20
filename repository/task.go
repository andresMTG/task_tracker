package repository

import (
	"time"
)

type Task struct {
	Id          int 
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	TODO = "todo"
	IN_PROGRESS = "in-progress"
	DONE = "done"
)