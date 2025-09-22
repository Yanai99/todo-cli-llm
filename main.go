package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Load tasks at startup
	if err := LoadTasks(); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <title>")
			return
		}
		title := os.Args[2]
		AddTask(title, "", "General", "No deadline")
		if err := SaveTasks(); err != nil {
			fmt.Println("Error saving tasks:", err)
		}

	case "list":
		ListTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <task ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		MarkTaskDoneById(id)
		if err := SaveTasks(); err != nil {
			fmt.Println("Error saving tasks:", err)
		}

	default:
		fmt.Println("Unknown command")
	}
}
