package types

import "time"

// Task is a definition of a task.
type Task struct {
	Name          string
	Note          string
	Flagged       bool
	DateAdded     time.Time
	DateCompleted time.Time
}
