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

func (l *TodoList) Complete(todo int) {
	if len(*l) < todo || todo < 1 {
		fmt.Errorf("Todo: %d was unable to be completed", todo)
	}

}
