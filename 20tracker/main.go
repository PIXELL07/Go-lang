package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const dataFile = "expenses.json"

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

func loadExpenses() []Expense {
	var expenses []Expense
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return expenses
	}
	json.Unmarshal(file, &expenses)
	return expenses
}

func saveExpenses(expenses []Expense) {
	data, _ := json.MarshalIndent(expenses, "", "  ")
	os.WriteFile(dataFile, data, 0644)
}

func getNextID(expenses []Expense) int {
	maxID := 0
	for _, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	return maxID + 1
}

func addExpense(description string, amount float64) {
	if amount <= 0 {
		fmt.Println("Amount must be greater than zero.")
		return
	}

	expenses := loadExpenses()
	id := getNextID(expenses)

	expense := Expense{
		ID:          id,
		Date:        time.Now().Format("2006-01-02"),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, expense)
	saveExpenses(expenses)
	fmt.Printf("Expense added successfully (ID: %d)\n", id)
}

func listExpenses() {
	expenses := loadExpenses()
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Println("ID  Date        Description     Amount")
	for _, e := range expenses {
		fmt.Printf("%-3d %-10s %-15s $%.2f\n", e.ID, e.Date, e.Description, e.Amount)
	}
}

func deleteExpense(id int) {
	expenses := loadExpenses()
	newExpenses := []Expense{}
	found := false

	for _, e := range expenses {
		if e.ID == id {
			found = true
			continue
		}
		newExpenses = append(newExpenses, e)
	}

	if !found {
		fmt.Println("Expense not found.")
		return
	}

	saveExpenses(newExpenses)
	fmt.Println("Expense deleted successfully.")
}

func updateExpense(id int, desc string, amount float64) {
	expenses := loadExpenses()
	found := false

	for i, e := range expenses {
		if e.ID == id {
			if desc != "" {
				expenses[i].Description = desc
			}
			if amount > 0 {
				expenses[i].Amount = amount
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Expense not found.")
		return
	}

	saveExpenses(expenses)
	fmt.Println("Expense updated successfully.")
}

func summary(month int) {
	expenses := loadExpenses()
	total := 0.0
	currentYear := time.Now().Year()

	for _, e := range expenses {
		date, _ := time.Parse("2006-01-02", e.Date)
		if month > 0 {
			if date.Month() == time.Month(month) && date.Year() == currentYear {
				total += e.Amount
			}
		} else {
			total += e.Amount
		}
	}

	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense-tracker [add|update|delete|list|summary]")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := addCmd.String("description", "", "Description")
		amount := addCmd.Float64("amount", 0, "Amount")
		addCmd.Parse(os.Args[2:])
		addExpense(*desc, *amount)

	case "list":
		listExpenses()

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.Int("id", 0, "Expense ID")
		delCmd.Parse(os.Args[2:])
		deleteExpense(*id)

	case "update":
		updCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updCmd.Int("id", 0, "Expense ID")
		desc := updCmd.String("description", "", "Description")
		amount := updCmd.Float64("amount", 0, "Amount")
		updCmd.Parse(os.Args[2:])
		updateExpense(*id, *desc, *amount)

	case "summary":
		sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := sumCmd.String("month", "", "Month")
		sumCmd.Parse(os.Args[2:])

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

	default:
		fmt.Println("Unknown command.")
	}
}
