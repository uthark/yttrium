package repo

import (
	"fmt"
	"time"

	"bitbucket.org/uthark/yttrium/internal/mongo"
	"bitbucket.org/uthark/yttrium/internal/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

const (
	// TaskCollection is a name of collection for storing tasks.
	TaskCollection string = "tasks"

	//LimitDefault specifies default limit for FindAll operation
	LimitDefault int = 1000
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
func (t TaskRepositoryImpl) Save(data *types.Task) (*types.Task, error) {

	session, collection := t.getTaskCollection()
	defer session.Close()

	if data.ID == "" {
		logger.Println("Saving new task", data)

		// new data
		// populating essentials
		data.DateAdded = time.Now()
		ID := uuid.New().String()
		data.ID = ID
		logger.Println("Assigning ID to task:", data.ID)

		err := collection.Insert(data)
		return data, err
	}

	logger.Println("Updating existing task", *data)

	err := collection.Update(byID(data.ID), data)
	return data, err

}

// List lists all tasks in a database.
func (t TaskRepositoryImpl) List() ([]*types.Task, error) {
	logger.Println("Getting all tasks.")
	session, collection := t.getTaskCollection()
	defer session.Close()

	var result []*types.Task
	q := bson.M{}
	err := collection.Find(q).Limit(LimitDefault).All(&result)
	return result, err
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
		panic(fmt.Errorf("can't connect to mongo db: %v", err))
	}
	session.SetMode(mgo.Primary, false)

	return TaskRepositoryImpl{
		dialInfo: dialInfo,
		session:  session,
	}
}

func byID(id string) bson.M {
	return bson.M{"_id": id}
}

func (t TaskRepositoryImpl) getTaskCollection() (session *mgo.Session, collection *mgo.Collection) {
	session = t.session.Copy()

	// use database provided in connection URL to connect.
	db := session.DB("")

	collection = db.C(TaskCollection)
	return session, collection
}
