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
		printUsage()
		return
	}
	command := os.Args[1]
	// Assume loadTasks() is defined elsewhere
	//tasks := loadTasks()
	var tasks []Task // Placeholder since loadTasks isn't defined

	switch command {
	case "add":
		if len(os.Args) < 3 {
			exitWith("Description required")
		}
		addTask(tasks, os.Args[2])

	case "update":
		requireArgs(4)
		updateTask(tasks, os.Args[2], os.Args[3])

	case "delete":
		requireArgs(3)
		deleteTask(tasks, os.Args[2])
	}
}

func printUsage() {
	panic("unimplemented")
}

func requireArgs(i int) {
	panic("unimplemented")
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
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func saveTasks(tasks []Task) {
	panic("unimplemented")
}

func updateTask(tasks []Task, idStr, desc string) {
	id := parseID(idStr)
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = now()
			saveTasks(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}
	exitWith("Task not found")
}

func now() string {
	panic("unimplemented")
}

func exitWith(s string) {
	panic("unimplemented")
}

func parseID(idStr string) any {
	panic("unimplemented")
}

func deleteTask(tasks []Task, idStr string) {
	id := parseID(idStr)
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Println("Task deleted successfully")
			return
		}
	}
	exitWith("Task not found")
}

func updateStatus(tasks []Task, idStr, status string) {
	id := parseID(idStr)
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = now()
			saveTasks(tasks)
			fmt.Printf("Task marked as %s\n", status)
			return
		}
	}
	exitWith("Task not found")
}
