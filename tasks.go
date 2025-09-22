package main

import (
	"fmt"
)

// Task represents a todo item
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Deadline    string `json:"deadline"` // can be time.Time later
	Done        bool   `json:"done"`
}

// Slice to hold tasks in memory
var tasks []Task

// AddTask adds a new task to the slice
func AddTask(title, description, category, deadline string) {
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Category:    category,
		Deadline:    deadline,
		Done:        false,
	}
	tasks = append(tasks, task)
	fmt.Printf("Task added:\nID: %d\nTitle: %s\n", task.ID, task.Title)
}

// ListTasks prints all tasks
func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s] - %s %s\n", t.ID, t.Title, t.Category, t.Deadline, status)
	}
}

// MarkTaskDone marks a task as done by ID
func MarkTaskDoneById(id int) {
	for i, t := range tasks {
		if t.ID == id {
			if t.Done {
				fmt.Println("Task already marked as done.")
				return
			}
			tasks[i].Done = true
			fmt.Printf("Task %d marked as done!\n", id)
			return
		}
	}
	fmt.Println("Task not found.")
}
