package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		printUsage()
		return
	}

	command := os.Args[1]
	tasks := loadTasks()

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

	case "mark-in-progress":
		requireArgs(3)
		updateStatus(tasks, os.Args[2], "in-progress")

	case "mark-done":
		requireArgs(3)
		updateStatus(tasks, os.Args[2], "done")

	case "list":
		if len(os.Args) == 3 {
			listTasks(tasks, os.Args[2])
		} else {
			listTasks(tasks, "")
		}

	default:
		printUsage()
	}
}

//----- Core Functions ------

// These functions modify the tasks slice in place and save to file immediately.

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

// TODO: Refactor to return updated tasks instead of modifying in place
// This will make it easier to test and avoid side effects

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

func listTasks(tasks []Task, filter string) {
	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf(
				"[#%d] %s (%s)\n",
				t.ID, t.Description, t.Status,
			)
		}
	}
}

// ----- Utilities ------

func loadTasks() []Task {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		exitWith("Failed to read tasks file")
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func parseID(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		exitWith("Invalid task ID")
	}
	return id
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func requireArgs(n int) {
	if len(os.Args) < n {
		printUsage()
		os.Exit(1)
	}
}

func exitWith(msg string) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func printUsage() {
	fmt.Println("Usage:")
	// ('go run task-cli.go <command> [args]')

	// Task Tracker CLI

	// Commands:
	//   add "description"
	//   update <id> "description"
	//   delete <id>
	//   mark-in-progress <id>
	//   mark-done <id>
	//   list
	//   list todo
	//   list in-progress
	//   list done

}
