package lib

import "time"

type TodoPriority string
type TodoStatus string

const (
	PriorityLow TodoPriority = "low"
	PriorityMedium TodoPriority = "medium"
	PriorityHigh TodoPriority = "high"
)

const (
	StatusPending TodoStatus = "pending"
	StatusInProgress TodoStatus = "in-progress"
	StatusDone TodoStatus = "done"
)

type Todo struct {
	name string
	status TodoStatus
	priority TodoPriority
	addedOn time.Time
	dueOn time.Time
}

func (todo Todo) print() {

}