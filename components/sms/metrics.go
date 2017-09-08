package sms

import (
	"github.com/kihamo/snitch"
)

const (
	MetricBalance   = ComponentName + "_balance_rubles_total"
	MetricTotalSend = ComponentName + "_send_total"
)

var (
	metricBalance   snitch.Gauge
	metricTotalSend snitch.Counter
)

type metricsCollector struct {
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	metricBalance.Describe(ch)
	metricTotalSend.Describe(ch)
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	metricBalance.Collect(ch)
	metricTotalSend.Collect(ch)
}

func (c *Component) Metrics() snitch.Collector {
	metricBalance = snitch.NewGauge(MetricBalance, "SMS balance in rubles")
	metricTotalSend = snitch.NewCounter(MetricTotalSend, "Number SMS sent")

	return &metricsCollector{}
}
