package usecases

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandManagerRoutesCommands(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"task_tracker", "add", "Managed task"}
	CommandManager("add", path)
	added := FileToListTask(path)
	if len(added) != 1 || added[0].Description != "Managed task" {
		t.Fatalf("expected add command to persist task, got %#v", added)
	}

	os.Args = []string{"task_tracker", "update", "1", "Managed update"}
	CommandManager("update", path)
	updated := FileToListTask(path)
	if updated[0].Description != "Managed update" {
		t.Fatalf("expected update command to change description, got %q", updated[0].Description)
	}

	os.Args = []string{"task_tracker", "list"}
	output := captureStdout(func() {
		CommandManager("list", path)
	})
	if !strings.Contains(output, "Managed update") {
		t.Fatalf("expected list command output to include task, got %q", output)
	}

	os.Args = []string{"task_tracker"}
	output = captureStdout(func() {
		CommandManager("help", path)
	})
	if !strings.Contains(output, "USAGE:") {
		t.Fatalf("expected help command output, got %q", output)
	}
}
