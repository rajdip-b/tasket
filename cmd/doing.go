package cmd

import (
	"fmt"
	"raj/tasket/lib/file"
	"raj/tasket/lib/todo"
	"strconv"

	"github.com/spf13/cobra"
)

var DoingCmd = &cobra.Command{
	Use:   "doing",
	Short: "Mark a todo as in progress",
	Example: `
	tasket doing 1
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		lastDisplayedTodos := file.LoadLastDisplayedTodos()

		// Get the todo details from the last listed todos
		var todoToDelete todo.Todo
		for i, todo := range lastDisplayedTodos.Todos {
			if i+1 == index {
				todoToDelete = todo
				break
			}
		}

		// Check if the todo exists
		if todoToDelete.Task == "" {
			fmt.Println("Todo not found")
			return
		}

		// Search for the todo in the todos file
		todos := file.LoadTodos()
		for i, t := range todos.Todos {
			if t.Equals(todoToDelete) {
				todos.Todos[i].Status = todo.StatusInProgress
				break
			}
		}

		file.WriteTodos(todos)
		
		fmt.Println("✅ Todo marked as in progress")
	},
}
