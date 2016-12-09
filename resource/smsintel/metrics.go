package smsintel

import (
	kit "github.com/go-kit/kit/metrics"
	"github.com/kihamo/shadow/resource/metrics"
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

func (r *Resource) MetricsRegister(m *metrics.Resource) {
	metricBalance = m.NewGauge(MetricSmsBalance)

	metricTotalSend := m.NewCounter(MetricSmsTotalSend)
	metricTotalSendSuccess = metricTotalSend.With("result", "success")
	metricTotalSendFailed = metricTotalSend.With("result", "failed")
}
