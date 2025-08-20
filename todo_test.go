package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	ls := todo.TodoList{}

	taskName := "Mow the lawn"

	ls.Add(taskName)

	if len(ls) == 0 || 1 < len(ls) {
		t.Errorf("Expected size %d, recieved size %d", 1, len(ls))
	}

	if ls[0].Task != taskName {
		t.Errorf("Expected task: %s, recieved %s", taskName, ls[0].Task)
	}
}

func TestComplete(t *testing.T) {
	ls := todo.TodoList{}

	taskName1 := "Go get milk"
	taskName2 := "Farm some minions"
	taskName3 := "Freeze wave"

	ls.Add(taskName1)
	ls.Add(taskName2)
	ls.Add(taskName3)

	if ls[0].Task != taskName1 {
		t.Errorf("Expected task: %s, recieved %s", taskName1, ls[0].Task)
	}

	ls.Complete(4)

	if len(ls) == 0 {
		t.Errorf("Completed unavailable task.")
	}

	ls.Complete(2)

	if !ls[1].Completed {
		t.Errorf("Todo %q was unable to be completed", ls[0].Task)
	}

	if ls[0].Completed {
		t.Errorf("Expected todo %d to be completed", 2)
	}
}

func TestDelete(t *testing.T) {
	ls := todo.TodoList{}

	task := "Go swimming"

	ls.Add(task)

	if len(ls) == 0 || ls[0].Task != task {
		t.Errorf("expected %q, got %q instead.", task, ls[0].Task)
	}

	ls.Delete(2)

	if len(ls) == 0 {
		t.Errorf("Deleted unavailable task.")
	}

	ls.Delete(1)

	if 0 < len(ls) {
		t.Errorf("Unable to delete task.")
	}
}

func TestSaveOpen(t *testing.T) {
	ls0 := todo.TodoList{}
	ls1 := todo.TodoList{}

	tasks := []string{
		"Go on a run",
		"Bathe",
		"Code",
	}

	for _, t := range tasks {
		ls0.Add(t)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Unable to create a temporary file.")
	}

	defer os.Remove(tf.Name())

	if err := ls0.Save(tf.Name()); err != nil {
		t.Fatalf("Unable to save list to temporary file %s.", err)
	}

	if err := ls1.Open(tf.Name()); err != nil {
		t.Fatalf("Unable to open temporary file")
	}

	for i := range ls0 {
		if ls0[i].Task != ls1[i].Task {
			t.Errorf("Expected %q, received %q.", ls0[i].Task, ls1[i].Task)
		}
	}
}
