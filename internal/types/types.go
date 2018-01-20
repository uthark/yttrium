package types

import "time"

// Task is a definition of a task.
type Task struct {
	// ID is a database id for a task.
	ID            string    `json:"id,omitempty" yaml:"id,omitempty" bson:"_id,omitempty"`
	Name          string    `json:"name" yaml:"name" bson:"name"`
	Note          string    `json:"note" yaml:"note" bson:"note"`
	Flagged       bool      `json:"flagged" yaml:"flagged" bson:"flagged"`
	DateAdded     time.Time `json:"dateAdded,omitempty" yaml:"dateAdded,omitempty" bson:"dateAdded"`
	DateCompleted time.Time `json:"dateCompleted,omitempty" yaml:"dateCompleted,omitempty" bson:"dateCompleted"`
}
