package todo_cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Task        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []Todo

func (l *TodoList) Add(taskName string) {
	todo := Todo{
		Task:        taskName,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, todo)
}

func (l *TodoList) Complete(todo int) error {
	if len(*l) < todo || todo < 1 {
		return fmt.Errorf("Todo: %d was unable to be completed", todo)
	}

	ls := *l

	ls[todo-1].Completed = true
	ls[todo-1].CompletedAt = time.Now()

	return nil
}

func (l *TodoList) Delete(todo int) error {
	if len(*l) < todo || todo < 1 {
		return fmt.Errorf("Todo: %d was unable to be deleted", todo)
	}

	ls := *l

	*l = append(ls[:todo-1], ls[todo:]...)

	return nil
}

func (l *TodoList) Save(filename string) error {
	if file, err := json.Marshal(l); err != nil {
		return fmt.Errorf("unable to marshal todo list")
	} else {
		return os.WriteFile(filename, file, 0644)
	}
}

func (l *TodoList) Open(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("unable to read file")
	}

	if err := json.Unmarshal(file, l); err != nil {
		return fmt.Errorf("failed to parse JSON from %q: %w", filename, err)
	}

	return nil
}
