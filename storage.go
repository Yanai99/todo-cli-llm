package main

import (
	"encoding/json"
	"os"
)

// SaveTasks saves the current tasks slice to tasks.json
func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// LoadTasks loads tasks from tasks.json into the tasks slice
func LoadTasks() error {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		tasks = []Task{} // no file yet, start empty
		return nil
	}

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tasks)
}
