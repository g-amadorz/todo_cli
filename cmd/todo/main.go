package main

import (
	"flag"
	"fmt"
	"os"
	"todo_cli"
)

const fileName = ".todos.json"

func main() {
	list := flag.Bool("list", false, "List all todos")
	complete := flag.Int("complete", 0, "Complete todo")
	add := flag.String("add", "", "Add todo")

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
	case *complete != 0:
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
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided")
		os.Exit(1)
	}

}
