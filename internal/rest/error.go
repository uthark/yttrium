package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/emicklei/go-restful"
)

// recoveryHandler catches panics and logs them.
// Returns 500 Error to the caller.
func recoveryHandler(panicReason interface{}, httpWriter http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Panic: recover from panic situation: - %v\r\n", panicReason))
	// skip first stacktrace lines to get only useful lines.
	for i := 4; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("\t%s:%d\r\n", file, line))
	}
	logger.Println(buffer.String())

	setContentTypeApplicationJSON(httpWriter)
	httpWriter.WriteHeader(http.StatusInternalServerError)

	msg := restful.ServiceError{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprint(panicReason),
	}
	b, _ := json.Marshal(msg)
	httpWriter.Write(b)
}

func serviceErrorHandler(err restful.ServiceError, req *restful.Request, resp *restful.Response) {
	logger.Println(req.Request.Method, req.Request.URL, err.Code, err.Message)
	setContentTypeApplicationJSON(resp)
	resp.WriteHeader(err.Code)
}

func notFound(resp http.ResponseWriter, req *http.Request) {
	logger.Println(req.URL, http.StatusNotFound)
	setContentTypeApplicationJSON(resp)
	resp.WriteHeader(http.StatusNotFound)
}

func setContentTypeApplicationJSON(httpWriter http.ResponseWriter) {
	httpWriter.Header().Set("Content-Type", restful.MIME_JSON)
}
