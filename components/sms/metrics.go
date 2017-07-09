package sms

import (
	"github.com/kihamo/snitch"
)

const (
	MetricBalance   = ComponentName + "_balance_rubles_total"
	MetricTotalSend = ComponentName + "_send_total"
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
	metricBalance = snitch.NewGauge(MetricBalance, "SMS balance in rubles")
	metricTotalSendSuccess = snitch.NewCounter(MetricTotalSend, "Number SMS sent with success status", "status", "success")
	metricTotalSendFailed = snitch.NewCounter(MetricTotalSend, "Number SMS sent with failed status", "status", "failed")

	return &metricsCollector{}
}
