package main

import (
	"flag"
	"fmt"
	"os"
	"todo_cli"
)

const fileName = ".todos.json"

func main() {
	list := flag.Bool("list", false, "List all the Todos in a Todo List")
	complete := flag.Int("complete", 0, "Mark the given index as a complete Todo and remove from")
	add := flag.String("add", "", "Add a new Todo with the given task name")
	delete := flag.Int("delete", 0, "Remove the index provided from the Todo List")

	flag.Parse()

	ls := &todo_cli.TodoList{}

	if _, err := os.Stat(fileName); err != nil {
		os.Create(fileName)
	}

	if err := ls.Open(fileName); err != nil {
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
		if err := ls.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add != "":
		ls.Add(*add)
		if err := ls.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		ls.Delete(*delete)
		if err := ls.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided")
		os.Exit(1)
	}

}
