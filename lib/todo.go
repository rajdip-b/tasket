package lib

import (
	"fmt"
	"time"
)

type TodoPriority string
type TodoStatus string

const (
	PriorityLow    TodoPriority = "low"
	PriorityMedium TodoPriority = "medium"
	PriorityHigh   TodoPriority = "high"
)

const (
	StatusPending    TodoStatus = "pending"
	StatusInProgress TodoStatus = "in-progress"
	StatusDone       TodoStatus = "done"
)

type Todo struct {
	Name     string       `toml:"name"`
	Category  string       `toml:"category"`
	Status   TodoStatus   `toml:"status"`
	Priority TodoPriority `toml:"priority"`
	AddedOn  time.Time    `toml:"added_on"`
	DueOn    time.Time    `toml:"due_on"`
}

type TodoList struct {
	Todos []Todo `toml:"todos"`
}

func (todo Todo) String() string {
	return fmt.Sprintf("%s (%s)\t[%s]", todo.Name, todo.Category, todo.Priority)
}

func GetAll() TodoList {
	return LoadTodos()
}

func AddTodo(name, category string, priority TodoPriority) *Todo {
	todo := Todo{
		Name:     name,
		Category:  category,
		Status:   StatusPending,
		Priority: priority,
		AddedOn:  time.Now(),
		DueOn:    time.Now(),
	}

	todos := LoadTodos()
	todos.Todos = append(todos.Todos, todo)
	WriteTodos(todos)

	return &todo
}
