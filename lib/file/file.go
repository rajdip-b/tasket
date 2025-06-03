package file

import (
	"os"
	"path/filepath"
	"raj/tasket/lib/todo"

	"github.com/BurntSushi/toml"
)

func getUserHomeDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return userHomeDir
}

func getTodosFilePath() string {
	userHomeDir := getUserHomeDir()
	return filepath.Join(userHomeDir, ".tasket", "todos.toml")
}

func getLastDisplayedTodosFilePath() string {
	userHomeDir := getUserHomeDir()
	return filepath.Join(userHomeDir, ".tasket", "last_displayed_todos.toml")
}

func createTodosFile() {
	path := getTodosFilePath()
	var err error

	// Create the .tasket directory if it doesn't exist
	err = os.MkdirAll(getUserHomeDir()+"/.tasket", 0755)
	if err != nil {
		panic(err)
	}

	// Create the todos.txt file
	err = os.WriteFile(path, []byte(""), 0644)
	if err != nil {
		panic(err)
	}
}

func createLastDisplayedTodosFile() {
	path := getLastDisplayedTodosFilePath()
	var err error

	// Create the .tasket directory if it doesn't exist
	err = os.MkdirAll(getUserHomeDir()+"/.tasket", 0755)
	if err != nil {
		panic(err)
	}

	// Create the todos.txt file
	err = os.WriteFile(path, []byte(""), 0644)
	if err != nil {
		panic(err)
	}
}

func todosFileExists() bool {
	if _, err := os.Stat(getTodosFilePath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func LoadTodos() todo.TodoList {
	if !todosFileExists() {
		createTodosFile()
	}

	var todoList todo.TodoList

	_, err := toml.DecodeFile(getTodosFilePath(), &todoList)
	if err != nil {
		panic(err)
	}

	return todoList
}

func LoadLastDisplayedTodos() todo.TodoList {
	if !todosFileExists() {
		createTodosFile()
	}

	var lastDisplayedTodos todo.TodoList

	_, err := toml.DecodeFile(getLastDisplayedTodosFilePath(), &lastDisplayedTodos)
	if err != nil {
		panic(err)
	}

	return lastDisplayedTodos
}

func WriteLastDisplayedTodos(lastDisplayedTodos todo.TodoList) {
	f, _ := os.OpenFile(getLastDisplayedTodosFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()

	err := toml.NewEncoder(f).Encode(lastDisplayedTodos)
	if err != nil {
		panic(err)
	}
}

func WriteTodos(todoList todo.TodoList) {
	f, _ := os.OpenFile(getTodosFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()

	err := toml.NewEncoder(f).Encode(todoList)
	if err != nil {
		panic(err)
	}
}
