package rest

import (
	"fmt"
	"net/http"

	"github.com/uthark/yttrium/internal/mime"
	"github.com/uthark/yttrium/internal/task/api"
	"github.com/uthark/yttrium/internal/types"
	"github.com/emicklei/go-restful"
)

// TaskID is a constant for task-id
const TaskID = "task-id"

// NewService return definition of the metrics service.
func NewService() *restful.WebService {
	ws := new(restful.WebService)

	ws = ws.
		Path("/task").
		Doc("Endpoint to manipulate tasks.").
		Produces(restful.MIME_JSON, mime.MediaTypeApplicationYaml)

	rest := TaskREST{
		api: api.NewTaskAPI(),
	}

	getTasks := ws.GET("").To(rest.ListTasks).
		Doc("returns tasks.").
		Operation("listTasks")
	ws = ws.Route(getTasks)

	saveTask := ws.POST("").To(rest.SaveTask).
		Doc("save task").
		Operation("saveTask")
	ws = ws.Route(saveTask)

	deleteTask := ws.DELETE(fmt.Sprintf("/{%s}", TaskID)).To(rest.DeleteTask).
		Doc("delete task").
		Param(ws.PathParameter(TaskID, "Task ID")).
		Operation("deleteTask")
	ws = ws.Route(deleteTask)

	getTask := ws.GET(fmt.Sprintf("/{%s}", TaskID)).To(rest.GetTask).
		Doc("get task").
		Param(ws.PathParameter(TaskID, "Task ID")).
		Operation("getTask")
	ws = ws.Route(getTask)

	return ws
}

// TaskREST is a REST wrapper around API.
type TaskREST struct {
	api *api.TaskAPI
}

// ListTasks lists tasks.
func (t *TaskREST) ListTasks(req *restful.Request, resp *restful.Response) {

	tasks, err := t.api.ListTasks()
	if err != nil {
		logger.Println(err)
		err = resp.WriteHeaderAndEntity(http.StatusServiceUnavailable, restful.ServiceError{
			Code:    http.StatusServiceUnavailable,
			Message: "Unable to get list of tasks",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
		return
	}

	err = resp.WriteEntity(tasks)
	if err != nil {
		logger.Println("Unable to send response", err)
	}
}

// SaveTask saves task.
func (t *TaskREST) SaveTask(req *restful.Request, resp *restful.Response) {
	task := types.Task{}
	err := req.ReadEntity(&task)
	if err != nil {
		logger.Println(err)
		err = resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to read request",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
		return
	}

	savedTask, err := t.api.SaveTask(task)
	if err != nil {
		logger.Println(err)
		err = resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to save task",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
		return
	}

	err = resp.WriteEntity(savedTask)
	if err != nil {
		logger.Println("Unable to send response", err)
	}
}

// GetTask returns task with the given id.
func (t *TaskREST) GetTask(req *restful.Request, resp *restful.Response) {
	taskID := req.PathParameter(TaskID)
	if taskID == "" {
		err := resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to get task",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
	}

	task, err := t.api.GetTask(taskID)
	if err != nil {
		logger.Println(err)
		err = resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to delete task",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
		return
	}

	err = resp.WriteEntity(task)
	if err != nil {
		logger.Println("Unable to send response", err)
	}

}

// DeleteTask deletes task.
func (t *TaskREST) DeleteTask(req *restful.Request, resp *restful.Response) {
	taskID := req.PathParameter(TaskID)
	if taskID == "" {
		err := resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to delete task",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
	}

	err := t.api.DeleteTask(taskID)
	if err != nil {
		logger.Println(err)
		err = resp.WriteHeaderAndEntity(http.StatusBadRequest, restful.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Unable to delete task",
		})
		if err != nil {
			logger.Println("Unable to send response", err)
		}
		return
	}

	resp.WriteHeader(http.StatusOK)

}
