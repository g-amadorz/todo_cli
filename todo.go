package todo_cli

import (
	"fmt"
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
