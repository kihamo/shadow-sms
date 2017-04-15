package sms

import (
	"github.com/kihamo/snitch"
)

const (
	MetricBalance   = ComponentName + ".balance"
	MetricTotalSend = ComponentName + ".total_send"
)

var (
	metricBalance          snitch.Gauge
	metricTotalSendSuccess snitch.Counter
	metricTotalSendFailed  snitch.Counter
)

type metricsCollector struct {
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	ch <- metricBalance.Description()
	ch <- metricTotalSendSuccess.Description()
	ch <- metricTotalSendFailed.Description()
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	ch <- metricBalance
	ch <- metricTotalSendSuccess
	ch <- metricTotalSendFailed
}

func (c *Component) Metrics() snitch.Collector {
	metricBalance = snitch.NewGauge(MetricBalance)
	metricTotalSendSuccess = snitch.NewCounter(MetricTotalSend, "status", "success")
	metricTotalSendFailed = snitch.NewCounter(MetricTotalSend, "status", "failed")

	return &metricsCollector{}
}
