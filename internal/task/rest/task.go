package rest

import (
	"net/http"

	"bitbucket.org/uthark/yttrium/internal/mime"
	"bitbucket.org/uthark/yttrium/internal/task/api"
	"bitbucket.org/uthark/yttrium/internal/types"
	"github.com/emicklei/go-restful"
)

// NewService return definition of the metrics service.
func NewService() *restful.WebService {
	ws := new(restful.WebService)

	ws = ws.
		Path("/task").
		Doc("Endpoint to manipulate tasks.").
		Produces(mime.MediaTypeApplicationYaml, restful.MIME_JSON)

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
