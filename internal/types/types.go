package types

import "time"

// Task is a definition of a task.
type Task struct {
	// ID is a database id for a task.
	ID            string    `json:"id,omitempty" yaml:"id,omitempty" bson:"_id,omitempty"`
	Name          string    `json:"name" yaml:"name" bson:"name"`
	Note          string    `json:"note,omitempty" yaml:"note,omitempty" bson:"note"`
	Flagged       bool      `json:"flagged,omitempty" yaml:"flagged,omitempty" bson:"flagged"`
	DateAdded     time.Time `json:"dateAdded,omitempty" yaml:"dateAdded" bson:"dateAdded"`
	DateCompleted time.Time `json:"dateCompleted,omitempty" yaml:"dateCompleted" bson:"dateCompleted"`
}
