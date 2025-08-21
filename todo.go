// Package todo..
package todo

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

type TodoList struct {
	Name  string
	Todos []Todo
}

//type TodoList []Todo

func NewTodoList(todoListName string) TodoList {
	todoList := TodoList{
		Name:  todoListName,
		Todos: []Todo{},
	}
	return todoList
}

func (l *TodoList) Add(taskName string) {
	todo := Todo{
		Task:        taskName,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	l.Todos = append(l.Todos, todo)
}

func (l *TodoList) Complete(todo int) error {
	if len(l.Todos) < todo || todo < 1 {
		return fmt.Errorf("Todo: %d was unable to be completed", todo)
	}

	ls := l.Todos

	ls[todo-1].Completed = true
	ls[todo-1].CompletedAt = time.Now()

	return nil
}

func (l *TodoList) Delete(todo int) error {
	if len(l.Todos) < todo || todo < 1 {
		return fmt.Errorf("Todo: %d was unable to be deleted", todo)
	}

	ls := l.Todos

	l.Todos = append(ls[:todo-1], ls[todo:]...)

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

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *TodoList) List() {
	fmt.Println(l.Name)

	for i, item := range l.Todos {
		check := "☐"
		if item.Completed {
			check = "☑"
		}
		fmt.Printf("%d: %s %s\n", i+1, item.Task, check)
	}
}
