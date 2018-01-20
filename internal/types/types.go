package types

import "time"

// Task is a definition of a task.
type Task struct {
	Name          string    `json:"name" yaml:"name" bson:"name"`
	Note          string    `json:"note" yaml:"note" bson:"note"`
	Flagged       bool      `json:"flagged" yaml:"flagged" bson:"flagged"`
	DateAdded     time.Time `json:"dateAdded" yaml:"dateAdded" bson:"dateAdded"`
	DateCompleted time.Time `json:"dateCompleted" yaml:"dateCompleted" bson:"dateCompleted"`
}
