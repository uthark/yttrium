package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	c.Filter(updateMetrics)
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
