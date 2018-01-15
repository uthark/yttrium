package api

import "bitbucket.org/uthark/yttrium/internal/types"

type TaskAPI struct {
}

func NewTaskAPI() *TaskAPI {
	return &TaskAPI{}
}

func (api TaskAPI) SaveTask(t types.Task) (*types.Task, error) {
	return nil, nil
}

func (api TaskAPI) ListTasks() ([]*types.Task, error) {
	return nil, nil
}
