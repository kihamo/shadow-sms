package sms

import (
	kit "github.com/go-kit/kit/metrics"
	"github.com/kihamo/shadow/components/metrics"
)

const (
	MetricSmsBalance   = "sms.balance"
	MetricSmsTotalSend = "sms.total_send"
)

var (
	metricBalance          kit.Gauge
	metricTotalSendSuccess kit.Counter
	metricTotalSendFailed  kit.Counter
)

func (c *Component) MetricsRegister(m *metrics.Component) {
	metricBalance = m.NewGauge(MetricSmsBalance)

	metricTotalSend := m.NewCounter(MetricSmsTotalSend)
	metricTotalSendSuccess = metricTotalSend.With("result", "success")
	metricTotalSendFailed = metricTotalSend.With("result", "failed")
}
