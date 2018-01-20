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
	c.Filter(updateMetrics)
	c = c.Add(prom.NewService())
	c = c.Add(taskrest.NewService())

	address := fmt.Sprintf(":%d", config.DefaultConfiguration().HTTPPort)
	logger.Println("Staring listening on ", address)
	server := &http.Server{
		Addr:     address,
		Handler:  c,
		ErrorLog: logger,
	}

	s.setupSignalHandler(server)

	return server.ListenAndServe()
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
func (s Server) setupSignalHandler(server *http.Server) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	go func() {
		s := <-c
		logger.Println("Got signal:", s)
		if s == syscall.SIGTERM || s == syscall.SIGINT || s == os.Interrupt {
			logger.Println("Safe Shutting down.")
			ctx := context.Background()
			err := server.Shutdown(ctx)
			if err != nil {
				logger.Println("Failed to shutdown server: ", err)
			}
		} else if s == syscall.SIGKILL {
			logger.Println("Closing server.")
			err := server.Close()
			if err != nil {
				logger.Println("Failed to forcibly close server: ", err)
			}
		}
		os.Exit(2)
	}()
}
