package usecases

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/andresMTG/task_tracker/repository"
)

func TestAutoIncrementId(t *testing.T) {
	tasks := []*repository.Task{{Id: 1}, {Id: 3}, {Id: 2}}

	if got := autoIncrementId(tasks); got != 4 {
		t.Fatalf("expected next id to be 4, got %d", got)
	}

	if got := autoIncrementId(nil); got != 1 {
		t.Fatalf("expected next id for empty list to be 1, got %d", got)
	}
}

func TestGetTaskById(t *testing.T) {
	tasks := []*repository.Task{{Id: 7, Description: "Example"}}

	task, index, err := getTaskById(7, tasks)
	if err != nil {
		t.Fatalf("expected task to be found: %v", err)
	}
	if task.Description != "Example" {
		t.Fatalf("expected description %q, got %q", "Example", task.Description)
	}
	if index != 0 {
		t.Fatalf("expected index 0, got %d", index)
	}

	_, _, err = getTaskById(99, tasks)
	if err == nil {
		t.Fatal("expected error for missing task id")
	}
}

func TestCreateTaskPersistsTaskAndDefaults(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	CreateTask("Write tests", nil, path)
	loaded := FileToListTask(path)

	if len(loaded) != 1 {
		t.Fatalf("expected one task to be persisted, got %d", len(loaded))
	}
	if loaded[0].Status != repository.TODO {
		t.Fatalf("expected status %q, got %q", repository.TODO, loaded[0].Status)
	}
	if loaded[0].CreatedAt.IsZero() || loaded[0].UpdatedAt.IsZero() {
		t.Fatal("expected timestamps to be initialized")
	}
}

func TestUpdateDescriptionUpdatesExistingTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initial := []*repository.Task{{
		Id:          1,
		Description: "Old",
		Status:      repository.TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}}

	WriteFile(path, initial)
	loaded := FileToListTask(path)

	UpdateDescription(loaded, 1, "New", path)
	updated := FileToListTask(path)

	if updated[0].Description != "New" {
		t.Fatalf("expected description to be updated, got %q", updated[0].Description)
	}
	if !updated[0].UpdatedAt.After(updated[0].CreatedAt) {
		t.Fatal("expected UpdatedAt to be refreshed after update")
	}
}

func TestDeleteTaskRemovesTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initial := []*repository.Task{
		{Id: 1, Description: "First", Status: repository.TODO, CreatedAt: now, UpdatedAt: now},
		{Id: 2, Description: "Second", Status: repository.DONE, CreatedAt: now, UpdatedAt: now},
	}

	WriteFile(path, initial)
	loaded := FileToListTask(path)

	DeleteTask(loaded, 2, path)
	updated := FileToListTask(path)

	if len(updated) != 1 || updated[0].Id != 1 {
		t.Fatalf("expected only task 1 to remain, got %#v", updated)
	}
}

func TestUpdateStatusChangesStatus(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initial := []*repository.Task{{
		Id:          1,
		Description: "Task",
		Status:      repository.TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}}

	WriteFile(path, initial)
	loaded := FileToListTask(path)

	UpdateStatus(loaded, 1, "mark-in-progress", path)
	updated := FileToListTask(path)
	if updated[0].Status != repository.IN_PROGRESS {
		t.Fatalf("expected status %q, got %q", repository.IN_PROGRESS, updated[0].Status)
	}

	UpdateStatus(loaded, 1, "mark-in-done", path)
	updated = FileToListTask(path)
	if updated[0].Status != repository.DONE {
		t.Fatalf("expected status %q, got %q", repository.DONE, updated[0].Status)
	}
}

func TestFilterTasksAndShowTasks(t *testing.T) {
	tasks := []*repository.Task{
		{Id: 1, Description: "Todo item", Status: repository.TODO},
		{Id: 2, Description: "Done item", Status: repository.DONE},
	}

	filtered := filterTasks(repository.DONE, nil, tasks)
	if len(filtered) != 1 || filtered[0].Description != "Done item" {
		t.Fatalf("expected only done task to be returned, got %#v", filtered)
	}

	output := captureStdout(func() {
		ShowTasks(tasks, repository.DONE)
	})
	if !strings.Contains(output, "Done item") {
		t.Fatalf("expected output to contain done item, got %q", output)
	}

	output = captureStdout(func() {
		ShowTasks(tasks, "invalid")
	})
	if !strings.Contains(output, "not a valid parameter") {
		t.Fatalf("expected invalid filter message, got %q", output)
	}
}
