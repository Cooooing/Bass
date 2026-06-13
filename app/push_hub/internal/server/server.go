package server

import (
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// ServerProviderSet 是 server 层依赖集合。
var ServerProviderSet = wire.NewSet(
	NewGRPCServer,
)

var (
	_metricRequests metric.Int64Counter
	_metricSeconds  metric.Float64Histogram
)

func InitMetrics(name string) {
	meter := otel.Meter(name)
	_metricRequests, _ = metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	_metricSeconds, _ = metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
}
