package rest

import (
	"fmt"
	"net/http"

	"bitbucket.org/uthark/yttrium/internal/config"
	"bitbucket.org/uthark/yttrium/internal/mime"
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
	logger.Println("Configuring HTTP server.")
	restful.SetLogger(logger)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.RegisterEntityAccessor(mime.MediaTypeApplicationYaml, NewYamlReaderWriter(mime.MediaTypeApplicationYaml))

	c := restful.NewContainer()
	c.ServeMux = http.NewServeMux()

	c.DoNotRecover(false)
	c.RecoverHandler(recoveryHandler)
	c.ServiceErrorHandler(serviceErrorHandler)
	c.Handle("/", http.HandlerFunc(notFound))
	c = c.Add(prom.NewService())

	address := fmt.Sprintf(":%d", config.DefaultConfiguration().HTTPPort)
	logger.Println("Staring listening on ", address)
	server := &http.Server{
		Addr:     address,
		Handler:  c,
		ErrorLog: logger,
	}

	return server.ListenAndServe()
}
