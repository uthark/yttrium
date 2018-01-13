package prom

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewService return definition of the metrics service.
func NewService() *restful.WebService {
	ws := new(restful.WebService)

	ws = ws.
		Path("/metrics").
		Doc("Returns prometheus metrics").
		Produces("text/plain")

	api := NewMetricsREST()

	getMetrics := ws.GET("").To(api.GetMetrics).
		Doc("returns current metrics.").
		Operation("getMetrics")
	ws = ws.Route(getMetrics)

	return ws
}

// MetricsREST is an implementation of Metrics REST endpoint.
type MetricsREST struct {
	handler http.Handler
}

// NewMetricsREST creates new REST endpoint for prometheus metrics
func NewMetricsREST() *MetricsREST {
	options := promhttp.HandlerOpts{DisableCompression: true}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer, options)
	return &MetricsREST{
		handler: handler,
	}
}

// GetMetrics returns metrics.
func (c MetricsREST) GetMetrics(request *restful.Request, response *restful.Response) {
	c.handler.ServeHTTP(response, request.Request)
}

// HTTPRequestsDurationMilliseconds is a metrics for counting HTTP Requests processing duration.
var HTTPRequestsDurationMilliseconds = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_milliseconds",
		Help: "How many HTTP requests processed, partitioned by endpoint, HTTP method and status code",
		// Percentiles in milliseconds.
		Buckets: []float64{1, 3, 5, 10, 25, 50, 100, 200, 500, 750, 1000, 1500, 2000, 3000, 5000, 8000},
	},
	[]string{"endpoint", "method", "status"},
)

// HTTPRequestsTotal is a metrics for counting total HTTP Requests processed.
var HTTPRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, partitioned by endpoint, HTTP method and status code",
	},
	[]string{"endpoint", "method", "status"},
)

// init registers metrics.
func init() {
	prometheus.MustRegister(
		HTTPRequestsDurationMilliseconds,
		HTTPRequestsTotal,
	)
}
