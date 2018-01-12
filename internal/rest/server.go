package rest

import (
	"net/http"

	"bitbucket.org/uthark/yttrium/internal/prom"
	"github.com/emicklei/go-restful"
)

// Server is a HTTP server.
type Server struct {
	Port uint16
}

// NewServer creates new instance of server.
func NewServer() *Server {
	return &Server{}
}

// Start creates RESTful container and starts  accepting HTTP requests.
func (s *Server) Start() error {
	logger.Println("Starting server...")
	restful.SetLogger(logger)
	restful.DefaultResponseContentType(restful.MIME_JSON)

	c := restful.NewContainer()
	c.ServeMux = http.NewServeMux()

	c.DoNotRecover(false)
	c.RecoverHandler(recoveryHandler)
	c.ServiceErrorHandler(serviceErrorHandler)
	c.Handle("/", http.HandlerFunc(notFound))
	c = c.Add(prom.NewService())

	// TODO: Allow to override port.
	server := &http.Server{
		Addr:     ":8080",
		Handler:  c,
		ErrorLog: logger,
	}

	return server.ListenAndServe()
}
