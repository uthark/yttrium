package repo

import (
	"bitbucket.org/uthark/yttrium/internal/types"
)

// TaskRepository is a repository to access tasks.
type TaskRepository interface {
	// Save saves task to a database.
	Save(t *types.Task) (*types.Task, error)
	// List returns all tasks.
	List() ([]*types.Task, error)
}

// TaskRepositoryImpl is an implementation of task repository.
type TaskRepositoryImpl struct {
}

// Save saves a task into a database.
func (t TaskRepositoryImpl) Save(task *types.Task) (*types.Task, error) {
	return task, nil
}

// List lists all tasks in a database.
func (t TaskRepositoryImpl) List() ([]*types.Task, error) {
	return nil, nil
}

// NewTaskRepository creates new task repository.
func NewTaskRepository() TaskRepository {
	return TaskRepositoryImpl{}
}
