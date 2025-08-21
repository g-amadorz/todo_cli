package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const defaultFileName = ".todos"

func main() {
	list := flag.Bool("list", false, "List all the Todos in a Todo List")
	complete := flag.Int("complete", 0, "Mark the given index as a complete Todo and remove from")
	task := flag.String("task", "", "Add a new Todo with the given task name")
	delete := flag.Int("delete", 0, "Remove the index provided from the Todo List")
	fileName := flag.String("name", defaultFileName, "Provide filename for TodoList")

	flag.Parse()

	ls := todo.NewTodoList(*fileName)

	*fileName = *fileName + ".json"

	if _, err := os.Stat(*fileName); err != nil {
		os.Create(*fileName)
	}

	if err := ls.Open(*fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		ls.List()
	case *complete > 0:
		ls.Complete(*complete)
		if err := ls.Save(*fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		ls.Add(*task)
		if err := ls.Save(*fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		ls.Delete(*delete)
		if err := ls.Save(*fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided")
		os.Exit(1)
	}

}
