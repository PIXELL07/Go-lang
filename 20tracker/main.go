// Declare the main package (entry point of the Go program)
package main

// Import required standard library packages
import (
	"encoding/json" // Used for encoding/decoding JSON data
	"flag"          // Used for parsing command-line flags
	"fmt"           // Used for formatted input/output
	"os"            // Used for file system operations
	"strconv"       // Used for converting strings to integers
	"time"          // Used for working with dates and times
)

// Constant that stores the name of the JSON file where expenses are saved
const dataFile = "expenses.json"

// Expense struct defines the structure of an expense record
type Expense struct {
	ID          int     `json:"id"`          // Unique ID for the expense
	Date        string  `json:"date"`        // Date of the expense
	Description string  `json:"description"` // Description of the expense
	Amount      float64 `json:"amount"`      // Amount spent
}

// loadExpenses reads expenses from the JSON file
func loadExpenses() []Expense {
	var expenses []Expense // Slice to hold expenses

	// Read the JSON file
	file, err := os.ReadFile(dataFile)
	if err != nil {
		// If file doesn't exist or can't be read, return empty slice
		return expenses
	}

	// Convert JSON data into Go struct
	json.Unmarshal(file, &expenses)

	// Return the loaded expenses
	return expenses
}

// saveExpenses writes the expense list to the JSON file
func saveExpenses(expenses []Expense) {
	// Convert expense slice to formatted JSON
	data, _ := json.MarshalIndent(expenses, "", "  ")

	// Write JSON data to file with read/write permissions
	os.WriteFile(dataFile, data, 0644)
}

// getNextID generates the next available expense ID
func getNextID(expenses []Expense) int {
	maxID := 0 // Track the highest ID found

	// Loop through expenses to find the max ID
	for _, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
		}
	}

	// Return next ID
	return maxID + 1
}

// addExpense adds a new expense to the file
func addExpense(description string, amount float64) {
	// Validate amount
	if amount <= 0 {
		fmt.Println("Amount must be greater than zero.")
		return
	}

	// Load existing expenses
	expenses := loadExpenses()

	// Generate new ID
	id := getNextID(expenses)

	// Create new expense object
	expense := Expense{
		ID:          id,
		Date:        time.Now().Format("2006-01-02"), // Current date
		Description: description,
		Amount:      amount,
	}

	// Append new expense
	expenses = append(expenses, expense)

	// Save updated expenses
	saveExpenses(expenses)

	// Print success message
	fmt.Printf("Expense added successfully (ID: %d)\n", id)
}

// listExpenses displays all saved expenses
func listExpenses() {
	// Load expenses from file
	expenses := loadExpenses()

	// If no expenses exist
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	// Print table header
	fmt.Println("ID  Date        Description     Amount")

	// Loop through expenses and print each one
	for _, e := range expenses {
		fmt.Printf("%-3d %-10s %-15s $%.2f\n",
			e.ID, e.Date, e.Description, e.Amount)
	}
}

// deleteExpense removes an expense by ID
func deleteExpense(id int) {
	// Load expenses
	expenses := loadExpenses()

	// Slice to store remaining expenses
	newExpenses := []Expense{}

	found := false // Track if expense was found

	// Loop through expenses
	for _, e := range expenses {
		if e.ID == id {
			found = true // Skip the matching expense
			continue
		}
		newExpenses = append(newExpenses, e)
	}

	// If expense was not found
	if !found {
		fmt.Println("Expense not found.")
		return
	}

	// Save updated expenses
	saveExpenses(newExpenses)

	fmt.Println("Expense deleted successfully.")
}

// updateExpense updates an existing expense
func updateExpense(id int, desc string, amount float64) {
	// Load expenses
	expenses := loadExpenses()

	found := false // Track if expense exists

	// Loop through expenses
	for i, e := range expenses {
		if e.ID == id {
			// Update description if provided
			if desc != "" {
				expenses[i].Description = desc
			}

			// Update amount if valid
			if amount > 0 {
				expenses[i].Amount = amount
			}

			found = true
			break
		}
	}

	// If expense not found
	if !found {
		fmt.Println("Expense not found.")
		return
	}

	// Save updated expenses
	saveExpenses(expenses)

	fmt.Println("Expense updated successfully.")
}

// summary calculates total expenses
func summary(month int) {
	// Load expenses
	expenses := loadExpenses()

	total := 0.0 // Store total expense amount
	currentYear := time.Now().Year()

	// Loop through expenses
	for _, e := range expenses {
		// Convert string date to time.Time
		date, _ := time.Parse("2006-01-02", e.Date)

		// If month filter is provided
		if month > 0 {
			if date.Month() == time.Month(month) && date.Year() == currentYear {
				total += e.Amount
			}
		} else {
			// Otherwise include all expenses
			total += e.Amount
		}
	}

	//  Print summary
	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

// main function is the entry point of the program
func main() {
	// Ensure a command is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense-tracker [add|update|delete|list|summary]")
		return
	}

	// Switch based on the command
	switch os.Args[1] {

	// ADD COMMAND
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := addCmd.String("description", "", "Description")
		amount := addCmd.Float64("amount", 0, "Amount")
		addCmd.Parse(os.Args[2:])
		addExpense(*desc, *amount)

	// LIST COMMAND
	case "list":
		listExpenses()

	// DELETE COMMAND
	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.Int("id", 0, "Expense ID")
		delCmd.Parse(os.Args[2:])
		deleteExpense(*id)

	// UPDATE COMMAND
	case "update":
		updCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updCmd.Int("id", 0, "Expense ID")
		desc := updCmd.String("description", "", "Description")
		amount := updCmd.Float64("amount", 0, "Amount")
		updCmd.Parse(os.Args[2:])
		updateExpense(*id, *desc, *amount)

	// SUMMARY COMMAND
	case "summary":
		sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := sumCmd.String("month", "", "Month")
		sumCmd.Parse(os.Args[2:])

		// If month is provided
		if *month != "" {
			m, err := strconv.Atoi(*month)
			if err != nil || m < 1 || m > 12 {
				fmt.Println("Invalid month.")
				return
			}
			summary(m)
		} else {
			summary(0)
		}

	// UNKNOWN COMMAND
	default:
		fmt.Println("Unknown command.")
	}
}
