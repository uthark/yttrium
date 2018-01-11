package rest

import (
	"log"
	"net/http"

	"os"

	"github.com/emicklei/go-restful"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

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
	logger.Println("Staring server...")
	c := restful.NewContainer()

	// TODO: Allow to override port.
	server := &http.Server{Addr: ":8080", Handler: c}

	return server.ListenAndServe()

}
