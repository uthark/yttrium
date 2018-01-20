package repo

import (
	"fmt"
	"time"

	"bitbucket.org/uthark/yttrium/internal/mongo"
	"bitbucket.org/uthark/yttrium/internal/types"
	"gopkg.in/mgo.v2"
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
	dialInfo *mgo.DialInfo
	session  *mgo.Session
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
func NewTaskRepository(mongoURL string) TaskRepository {
	dialInfo, err := mongo.ParseURL(mongoURL)
	if err != nil {
		panic(fmt.Sprintf("Invalid mongo db url: %v", err))
	}
	dialInfo.Timeout = 10 * time.Second
	dialInfo.FailFast = true
	dialInfo.PoolLimit = 50

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(fmt.Errorf("Can't connect to mongo db: %v", err))
	}
	session.SetMode(mgo.Primary, false)

	return TaskRepositoryImpl{
		dialInfo: dialInfo,
		session:  session,
	}
}
