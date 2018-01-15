package task

import (
	"bitbucket.org/uthark/yttrium/internal/api"
	"bitbucket.org/uthark/yttrium/internal/mime"
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

type TaskREST struct {
	api *api.TaskAPI
}

func (t *TaskREST) ListTasks(req *restful.Request, resp *restful.Response) {
	t.api.ListTasks()
}

func (t *TaskREST) SaveTask(req *restful.Request, resp *restful.Response) {
	// t.api.SaveTask()
}
