package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const defaultFileName = ".todos.json"

func main() {
	list := flag.Bool("list", false, "List all the Todos in a Todo List")
	complete := flag.Int("complete", 0, "Mark the given index as a complete Todo and remove from")
	task := flag.String("task", "", "Add a new Todo with the given task name")
	delete := flag.Int("delete", 0, "Remove the index provided from the Todo List")
	todoListName := flag.String("open", defaultFileName, "If provided open given todo list file")

	flag.Parse()

	ls := &todo.TodoList{}

	if _, err := os.Stat(*todoListName); err != nil {
		os.Create(defaultFileName)
	}

	if err := ls.Open(*todoListName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for i, item := range *ls {
			check := "☐"
			if item.Completed {
				check = "☑"
			}
			fmt.Printf("%d: %s %s\n", i+1, item.Task, check)
		}
	case *complete > 0:
		ls.Complete(*complete)
		if err := ls.Save(*todoListName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		ls.Add(*task)
		if err := ls.Save(*todoListName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		ls.Delete(*delete)
		if err := ls.Save(*todoListName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided")
		os.Exit(1)
	}

}
