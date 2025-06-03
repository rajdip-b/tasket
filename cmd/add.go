package cmd

import (
	"fmt"
	"os"
	"raj/tasket/lib/file"
	"raj/tasket/lib/todo"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [item to add]",
	Short: "Adds an item to the todo list",
	Example: `
	tasket add "Read Go documentation +golang @high"
	tasket add "Read Go documentation +golang"
	tasket add "Read Go documentation"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fullTaskString := strings.Join(args, "")

		category, categoryDirectivePosition := extractCategory(fullTaskString)
		priority, priorityDirectivePosition := extractPriority(fullTaskString)
		task := extractTask(fullTaskString, categoryDirectivePosition, priorityDirectivePosition)

		todo := todo.Todo{
			Task:     task,
			Category: category,
			Status:   todo.StatusPending,
			Priority: priority,
			AddedOn:  time.Now(),
			DueOn:    time.Now(),
		}

		todos := file.LoadTodos()
		todos.Todos = append(todos.Todos, todo)
		file.WriteTodos(todos)
		fmt.Println("Task Added: " + todo.String(true))
	},
}

func extractTask(fullTaskString string, categoryDirectivePosition int, priorityDirectivePosition int) string {
	var task string = ""

	if categoryDirectivePosition == -1 && priorityDirectivePosition == -1 {
		task = fullTaskString
	} else if priorityDirectivePosition == -1 {
		task = fullTaskString[:priorityDirectivePosition]
	} else if categoryDirectivePosition == -1 {
		task = fullTaskString[:categoryDirectivePosition]
	} else {
		min := min(categoryDirectivePosition, priorityDirectivePosition)
		task = fullTaskString[:min]
	}
	task = strings.Trim(task, " ")

	if task == "" {
		fmt.Println("No task specified")
		os.Exit(1)
	}

	return task
}

func extractCategory(fullTaskString string) (string, int) {
	var category string = "uncategorized"
	categoryDirectivePosition := strings.LastIndex(fullTaskString, "+")

	if categoryDirectivePosition != -1 {
		category = strings.Trim(strings.Split(fullTaskString[categoryDirectivePosition+1:], " ")[0], " ")
	}

	if category == "" {
		fmt.Println("No category specified")
		os.Exit(1)
	}

	return category, categoryDirectivePosition
}

func extractPriority(fullTaskString string) (todo.TodoPriority, int) {
	var priority string = "low"
	priorityDirectivePosition := strings.LastIndex(fullTaskString, "@")

	if priorityDirectivePosition != -1 {
		priorityString := strings.Trim(strings.Split(fullTaskString[priorityDirectivePosition+1:], " ")[0], " ")

		if priorityString != "low" && priorityString != "medium" && priorityString != "high" {
			fmt.Println("Invalid priority value specified: " + priorityString)
			os.Exit(1)
		}

		priority = priorityString
	}

	var priorityEnum todo.TodoPriority

	switch priority {
	case "low":
		priorityEnum = todo.PriorityLow
	case "medium":
		priorityEnum = todo.PriorityMedium
	case "high":
		priorityEnum = todo.PriorityHigh
	}

	return priorityEnum, priorityDirectivePosition
}
