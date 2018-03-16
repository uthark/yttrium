package api

import (
	"time"

	"github.com/uthark/yttrium/internal/config"
	"github.com/uthark/yttrium/internal/task/repo"
	"github.com/uthark/yttrium/internal/types"
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

// DeleteTask saves task.
func (api TaskAPI) DeleteTask(taskID string) error {
	return api.repo.Delete(taskID)
}

// GetTask saves task.
func (api TaskAPI) GetTask(taskID string) (*types.Task, error) {
	return api.repo.FindByID(taskID)
}
