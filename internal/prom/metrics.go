package prom

import (
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

	api := MetricsREST{}

	getMetrics := ws.GET("").To(api.GetMetrics).
		Doc("returns current metrics.").
		Operation("getMetrics")
	ws = ws.Route(getMetrics)

	return ws
}

// MetricsREST is an implementation of Metrics REST endpoint.
type MetricsREST struct{}

// GetMetrics returns metrics.
func (c MetricsREST) GetMetrics(request *restful.Request, response *restful.Response) {
	options := promhttp.HandlerOpts{DisableCompression: true}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer, options)
	handler.ServeHTTP(response, request.Request)
}
