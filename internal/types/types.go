package types

import "time"

// Task is a definition of a task.
type Task struct {
	Name          string    `json:"name" yaml:"name"`
	Note          string    `json:"note" yaml:"note"`
	Flagged       bool      `json:"flagged" yaml:"flagged"`
	DateAdded     time.Time `json:"dateAdded" yaml:"dateAdded"`
	DateCompleted time.Time `json:"dateCompleted" yaml:"dateCompleted"`
}
