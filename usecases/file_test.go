package usecases

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/andresMTG/task_tracker/repository"
)

func TestFileExists(t *testing.T) {
	t.Run("returns false for missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.json")
		if FileExists(path) {
			t.Fatalf("expected missing file to return false")
		}
	})

	t.Run("returns true for existing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "existing.json")
		file, err := os.Create(path)
		if err != nil {
			t.Fatalf("could not create temp file: %v", err)
		}
		defer file.Close()

		if !FileExists(path) {
			t.Fatalf("expected existing file to return true")
		}
	})
}

func TestCreateFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	file := CreateFile(path)
	if file == nil {
		t.Fatal("expected a file to be created")
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("expected file to exist: %v", err)
	}
	if info.Size() != 0 {
		t.Fatalf("expected new file to be empty, got size %d", info.Size())
	}

	if _, err := file.WriteString("old data"); err != nil {
		t.Fatalf("could not write initial content: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("could not close file: %v", err)
	}

	file = CreateFile(path)
	if file == nil {
		t.Fatal("expected file to be recreated")
	}
	defer file.Close()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("could not read recreated file: %v", err)
	}
	if string(content) != "" {
		t.Fatalf("expected file to be truncated, got %q", string(content))
	}
}

func TestWriteFileAndFileToListTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initialTasks := []*repository.Task{{
		Id:          1,
		Description: "Buy milk",
		Status:      repository.TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}}

	WriteFile(path, initialTasks)
	loadedTasks := FileToListTask(path)

	if len(loadedTasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Description != initialTasks[0].Description {
		t.Fatalf("expected description %q, got %q", initialTasks[0].Description, loadedTasks[0].Description)
	}
	if loadedTasks[0].Status != repository.TODO {
		t.Fatalf("expected status %q, got %q", repository.TODO, loadedTasks[0].Status)
	}
}

func TestCreateTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	CreateTask("Buy milk", nil, path)
	loadedTasks := FileToListTask(path)

	if len(loadedTasks) != 1 {
		t.Fatalf("expected 1 task to be created, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Id != 1 {
		t.Fatalf("expected first task id to be 1, got %d", loadedTasks[0].Id)
	}
	if loadedTasks[0].Description != "Buy milk" {
		t.Fatalf("expected description to be persisted, got %q", loadedTasks[0].Description)
	}
	if loadedTasks[0].Status != repository.TODO {
		t.Fatalf("expected default status %q, got %q", repository.TODO, loadedTasks[0].Status)
	}
	if loadedTasks[0].CreatedAt.IsZero() || loadedTasks[0].UpdatedAt.IsZero() {
		t.Fatalf("expected timestamps to be populated")
	}
}

func TestUpdateDescription(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initialTasks := []*repository.Task{{
		Id:          1,
		Description: "Old description",
		Status:      repository.TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}}

	WriteFile(path, initialTasks)
	loadedTasks := FileToListTask(path)
	UpdateDescription(loadedTasks, 1, "New description", path)

	updatedTasks := FileToListTask(path)
	if len(updatedTasks) != 1 {
		t.Fatalf("expected 1 task after update, got %d", len(updatedTasks))
	}
	if updatedTasks[0].Description != "New description" {
		t.Fatalf("expected updated description, got %q", updatedTasks[0].Description)
	}
	if updatedTasks[0].UpdatedAt.Before(updatedTasks[0].CreatedAt) {
		t.Fatalf("expected updated time to be refreshed")
	}
}

func TestDeleteTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initialTasks := []*repository.Task{
		{Id: 1, Description: "First", Status: repository.TODO, CreatedAt: now, UpdatedAt: now},
		{Id: 2, Description: "Second", Status: repository.DONE, CreatedAt: now, UpdatedAt: now},
	}

	WriteFile(path, initialTasks)
	loadedTasks := FileToListTask(path)
	DeleteTask(loadedTasks, 2, path)

	updatedTasks := FileToListTask(path)
	if len(updatedTasks) != 1 {
		t.Fatalf("expected 1 task after deletion, got %d", len(updatedTasks))
	}
	if updatedTasks[0].Id != 1 {
		t.Fatalf("expected remaining task id to be 1, got %d", updatedTasks[0].Id)
	}
}

func TestUpdateStatus(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	now := time.Now()
	initialTasks := []*repository.Task{{
		Id:          1,
		Description: "Task",
		Status:      repository.TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}}

	WriteFile(path, initialTasks)
	loadedTasks := FileToListTask(path)
	UpdateStatus(loadedTasks, 1, "mark-in-progress", path)

	inProgressTasks := FileToListTask(path)
	if inProgressTasks[0].Status != repository.IN_PROGRESS {
		t.Fatalf("expected status %q, got %q", repository.IN_PROGRESS, inProgressTasks[0].Status)
	}

	UpdateStatus(loadedTasks, 1, "mark-in-done", path)
	doneTasks := FileToListTask(path)
	if doneTasks[0].Status != repository.DONE {
		t.Fatalf("expected status %q, got %q", repository.DONE, doneTasks[0].Status)
	}
}

func TestShowTasks(t *testing.T) {
	tasks := []*repository.Task{
		{Id: 1, Description: "Todo task", Status: repository.TODO},
		{Id: 2, Description: "Done task", Status: repository.DONE},
	}

	output := captureStdout(func() {
		ShowTasks(tasks, repository.DONE)
	})
	if !strings.Contains(output, "Done task") {
		t.Fatalf("expected filtered output to contain done task, got %q", output)
	}
	if strings.Contains(output, "Todo task") {
		t.Fatalf("expected filtered output to exclude todo task")
	}

	output = captureStdout(func() {
		ShowTasks(tasks, "invalid")
	})
	if !strings.Contains(output, "not a valid parameter") {
		t.Fatalf("expected invalid filter message, got %q", output)
	}
}

func TestCommandManager(t *testing.T) {
	t.Run("add command persists a task", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "tasks.json")
		oldArgs := os.Args
		os.Args = []string{"task_tracker", "add", "New task"}
		t.Cleanup(func() { os.Args = oldArgs })

		CommandManager("add", path)
		loadedTasks := FileToListTask(path)
		if len(loadedTasks) != 1 {
			t.Fatalf("expected 1 task after add command, got %d", len(loadedTasks))
		}
		if loadedTasks[0].Description != "New task" {
			t.Fatalf("expected added task description, got %q", loadedTasks[0].Description)
		}
	})

	t.Run("list command prints stored tasks", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "tasks.json")
		oldArgs := os.Args
		os.Args = []string{"task_tracker", "list"}
		t.Cleanup(func() { os.Args = oldArgs })

		CreateTask("Persisted task", nil, path)
		output := captureStdout(func() {
			CommandManager("list", path)
		})
		if !strings.Contains(output, "Persisted task") {
			t.Fatalf("expected list output to contain task, got %q", output)
		}
	})

	t.Run("help is shown when no subcommand is provided", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "tasks.json")
		oldArgs := os.Args
		os.Args = []string{"task_tracker"}
		t.Cleanup(func() { os.Args = oldArgs })

		output := captureStdout(func() {
			CommandManager("add", path)
		})
		if !strings.Contains(output, "USAGE:") {
			t.Fatalf("expected help output, got %q", output)
		}
	})
}

func captureStdout(fn func()) string {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()
	if err := w.Close(); err != nil {
		panic(err)
	}
	output, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(output)
}
