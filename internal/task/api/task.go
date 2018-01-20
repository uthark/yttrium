package api

import (
	"time"

	"bitbucket.org/uthark/yttrium/internal/config"
	"bitbucket.org/uthark/yttrium/internal/task/repo"
	"bitbucket.org/uthark/yttrium/internal/types"
)

// TaskAPI is an API to work with tasks.
type TaskAPI struct {
	repo repo.TaskRepository
}

// NewTaskAPI creates new API to work with tasks.
func NewTaskAPI() *TaskAPI {
	return &TaskAPI{
		repo: repo.NewTaskRepository(config.DefaultConfiguration().DatabaseConnection),
	}
}

// SaveTask saves task.
func (api TaskAPI) SaveTask(t types.Task) (*types.Task, error) {
	if t.DateAdded.IsZero() {
		t.DateAdded = time.Now()
	}
	return api.repo.Save(&t)
}

// ListTasks lists task.
func (api TaskAPI) ListTasks() ([]*types.Task, error) {
	return api.repo.List()
}
