package main

import (
	"fmt"
	"os"
	"time"
)

const fileName = "tasks.json"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	if len(os.Args) < 2 {
		// Assume printUsage() is defined elsewhere
		// printUsage()
		return
	}
	command := os.Args[1]
	// Assume loadTasks() is defined elsewhere
	// tasks := loadTasks()
	var tasks []Task // Placeholder since loadTasks isn't defined

	switch command {
	case "add":
		if len(os.Args) < 3 {
			// Assume exitWith() is defined elsewhere
			// exitWith("Description required")
			fmt.Println("Description required")
			return
		}
		addTask(tasks, os.Args[2])
	}
}

// addTask function definition moved outside the main function
func addTask(tasks []Task, description string) {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	now := time.Now().UTC().Format(time.RFC3339)
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	// Assume saveTasks() is defined elsewhere
	// saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}
