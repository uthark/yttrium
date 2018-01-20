package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"bitbucket.org/uthark/yttrium/internal/config"
	"bitbucket.org/uthark/yttrium/internal/mime"
	"bitbucket.org/uthark/yttrium/internal/prom"
	taskrest "bitbucket.org/uthark/yttrium/internal/task/rest"
	"github.com/emicklei/go-restful"
)

// Server is a HTTP server.
type Server struct {
	Port   uint16
	server *http.Server
	stop   chan int
}

// NewServer creates new instance of server.
func NewServer() *Server {
	return &Server{}
}
func (s *Server) Restart() {
	logger.Println("Restarting server")
	s.Stop()
	go s.Start()
}

func (s *Server) Init(stop chan int) {
	logger.Println("Configuring HTTP server.")

	s.setupSignalHandler()
	s.stop = stop

	restful.SetLogger(logger)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.RegisterEntityAccessor(mime.MediaTypeApplicationYaml, NewYamlReaderWriter(mime.MediaTypeApplicationYaml))

}

func (s *Server) Stop() {
	logger.Println("Stopping server.")
	ctx := context.TODO()
	err := s.server.Shutdown(ctx)
	if err != nil {
		logger.Println(err)
	}
	logger.Println("Server stopped")
}

// Start creates RESTful container and starts  accepting HTTP requests.
func (s *Server) Start() {
	logger.Println("Starting server.")
	address := fmt.Sprintf(":%d", config.DefaultConfiguration().HTTPPort)
	logger.Println("Starting listening on ", address)

	c := restful.NewContainer()
	c.ServeMux = http.NewServeMux()

	c.DoNotRecover(false)
	c.RecoverHandler(recoveryHandler)
	c.ServiceErrorHandler(serviceErrorHandler)
	c.Handle("/", http.HandlerFunc(notFound))
	c.Filter(updateMetrics)
	c = c.Add(prom.NewService())
	c = c.Add(taskrest.NewService())

	server := &http.Server{
		Addr:     address,
		Handler:  c,
		ErrorLog: logger,
	}
	s.server = server

	err := server.ListenAndServe()
	if err != nil {
		logger.Println(err)
	}

}

// webserviceLogging logs requested HTTP URL and method
func webserviceLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Printf("Started [HTTP] %s %s\n", req.Request.Method, req.Request.URL)
	start := time.Now()
	chain.ProcessFilter(req, resp)
	logger.Printf("Finished [HTTP] %s %s in %dns \n", req.Request.Method, req.Request.URL, time.Now().Sub(start).Nanoseconds())
}

// updateMetrics updates metrics
func updateMetrics(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	startTime := time.Now()
	chain.ProcessFilter(req, resp)
	durationMilliseconds := time.Now().Sub(startTime).Nanoseconds() / int64(time.Millisecond)
	endpoint := req.SelectedRoutePath()
	method := req.Request.Method
	httpStatus := strconv.Itoa(resp.StatusCode())
	prom.HTTPRequestsTotal.WithLabelValues(endpoint, method, httpStatus).Inc()
	prom.HTTPRequestsDurationMilliseconds.WithLabelValues(endpoint, method, httpStatus).Observe(float64(durationMilliseconds))
}

// setupSignalHandler listens for syscall signals and stops HTTP server.
func (s *Server) setupSignalHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	go func() {
		sig := <-c
		logger.Println("Got signal:", sig)
		if sig == syscall.SIGTERM || sig == syscall.SIGINT || sig == os.Interrupt {
			logger.Println("Safe Shutting down.")
			ctx := context.Background()
			err := s.server.Shutdown(ctx)
			if err != nil {
				logger.Println("Failed to shutdown server: ", err)
			}
		} else if sig == syscall.SIGKILL {
			logger.Println("Closing server.")
			err := s.server.Close()
			if err != nil {
				logger.Println("Failed to forcibly close server: ", err)
			}
		}
		s.stop <- 1
	}()
}
