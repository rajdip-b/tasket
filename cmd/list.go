package cmd

import (
	"fmt"
	"raj/tasket/lib/file"
	"raj/tasket/lib/todo"
	"strings"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the todos",
	Example: `
	tasket list"
	tasket list --category golang"
	tasket list --category golang --priority high"
	tasket list --filter "some task"`,
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		priority, _ := cmd.Flags().GetString("priority")
		filter, _ := cmd.Flags().GetString("filter")
		status, _ := cmd.Flags().GetString("status")

		categoryBasedTodos := make(map[string][]todo.Todo)

		allTodos := file.LoadTodos()
		for _, todo := range allTodos.Todos {
			// Filter by categody
			if category != "" && category != todo.Category {
				continue
			}

			// Filter by priority
			if priority != "" && priority != todo.Priority.String() {
				continue
			}

			// Filter by task name
			if filter != "" && !strings.Contains(todo.Task, filter) {
				continue
			}

			// Filter by status
			if status != "" && status != todo.Status.String() {
				continue
			}

			categoryBasedTodos[todo.Category] = append(categoryBasedTodos[todo.Category], todo)
		}

		var lastDisplayedTodos todo.TodoList
		index := 1
		for category, todos := range categoryBasedTodos {
			fmt.Printf("\n\033[1mCategory: %s\033[0m\n", category)
			fmt.Println("----------------------------------------------------------")
			for _, todo := range todos {
				lastDisplayedTodos.Todos = append(lastDisplayedTodos.Todos, todo)
				fmt.Println(todo.String(index))
				index++
			}
			fmt.Println("----------------------------------------------------------")
		}

		file.WriteLastDisplayedTodos(lastDisplayedTodos)
	},
}

func init() {
	ListCmd.Flags().StringP("category", "c", "", "Filter by category")
	ListCmd.Flags().StringP("priority", "p", "", "Filter by priority")
	ListCmd.Flags().StringP("filter", "f", "", "Filter by task name")
	ListCmd.Flags().StringP("status", "s", "", "Filter by status")
}
