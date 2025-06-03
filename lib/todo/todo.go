package todo

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

var priorityName = map[TodoPriority]string{
	PriorityLow:    "low",
	PriorityMedium: "medium",
	PriorityHigh:   "high",
}

func (priority TodoPriority) String() string {
	return priorityName[priority]
}

const (
	StatusPending    TodoStatus = "pending"
	StatusInProgress TodoStatus = "in-progress"
	StatusDone       TodoStatus = "done"
)

var statusName = map[TodoStatus]string{
	StatusPending:    "pending",
	StatusInProgress: "in-progress",
	StatusDone:       "done",
}

func (status TodoStatus) String() string {
	return statusName[status]
}

type Todo struct {
	Task     string       `toml:"task"`
	Category string       `toml:"category"`
	Status   TodoStatus   `toml:"status"`
	Priority TodoPriority `toml:"priority"`
	AddedOn  time.Time    `toml:"added_on"`
	DueOn    time.Time    `toml:"due_on"`
}

type TodoList struct {
	Todos []Todo `toml:"todos"`
}

func (todo Todo) String(showCategory bool) string {
	if showCategory {
		return fmt.Sprintf("%s (%s)\t[%s]", todo.Task, todo.Category, todo.Priority)
	} else {
		return fmt.Sprintf("%s\t[%s]", todo.Task, todo.Priority)
	}
}

func (todo Todo) Equals(other Todo) bool {
	return (todo.Task == other.Task && todo.Category == other.Category && todo.Status == other.Status && todo.Priority == other.Priority && todo.AddedOn == other.AddedOn && todo.DueOn == other.DueOn)
}
