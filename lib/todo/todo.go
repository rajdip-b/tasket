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

func StatusSymbol(status TodoStatus) string {
    switch status {
    case StatusPending:
        return "⏳" // Hourglass for Todo
    case StatusInProgress:
        return "🚧" // Construction for In Progress
    case StatusDone:
        return "✅" // Checkmark for Done
    default:
        return "❓"
    }
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

func (todo Todo) Equals(other Todo) bool {
	return (todo.Task == other.Task && todo.Category == other.Category && todo.Status == other.Status && todo.Priority == other.Priority && todo.AddedOn == other.AddedOn && todo.DueOn == other.DueOn)
}

func (t Todo) String(number int) string {
    // Example:  1. [🚧 inprogress]  learn go         [low]   due: 2025-05-30
    return fmt.Sprintf("%2d. [%s %-11s] %-25s [%-4s] due: %s",
        number,
        StatusSymbol(TodoStatus(t.Status)), t.Status,
        t.Task,
        t.Priority,
        t.DueOn.Format("2006-01-02"),
    )
}