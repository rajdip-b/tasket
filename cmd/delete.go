package cmd

import (
	"fmt"
	"raj/tasket/lib/file"
	"raj/tasket/lib/todo"
	"strconv"

	"slices"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Example: `
	tasket delete 1
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
		for i, todo := range todos.Todos {
			if todo.Equals(todoToDelete) {
				todos.Todos = slices.Delete(todos.Todos, i, i+1)
				break
			}
		}

		file.WriteTodos(todos)
		
		fmt.Println("✅ Todo deleted")
	},
}
