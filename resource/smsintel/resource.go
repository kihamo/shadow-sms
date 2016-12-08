package smsintel

import (
	kit "github.com/go-kit/kit/metrics"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
	"github.com/rs/xlog"
)

type Resource struct {
	application *shadow.Application
	config      *config.Resource
	client      *smsintel.SmsIntel
	logger      xlog.Logger

	metricBalance          kit.Gauge
	metricTotalSendSuccess kit.Counter
	metricTotalSendFailed  kit.Counter
}

func (r *Resource) GetName() string {
	return "smsintel"
}

func (r *Resource) Init(a *shadow.Application) error {
	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}

	r.config = resourceConfig.(*config.Resource)

	resourceLogger, err := a.GetResource("logger")
	if err == nil {
		r.logger = resourceLogger.(*logger.Resource).Get(r.GetName())
	}

	r.application = a

	return nil
}

func (r *Resource) Run() error {
	resourceMetrics, err := r.application.GetResource("metrics")
	if err == nil {
		rMetrics := resourceMetrics.(*metrics.Resource)

		r.metricBalance = rMetrics.NewGauge(MetricSmsBalance)

		metricTotalSend := rMetrics.NewCounter(MetricSmsTotalSend)
		r.metricTotalSendSuccess = metricTotalSend.With("result", "success")
		r.metricTotalSendFailed = metricTotalSend.With("result", "failed")
	}

	return nil
}

func (r *Resource) GetClient() *smsintel.SmsIntel {
	if r.client == nil {
		r.client = smsintel.NewSmsIntel(r.config.GetString("sms.login"), r.config.GetString("sms.password"))
	}

	return r.client
}

func (r *Resource) Send(message, phone string) error {
	input := &procedure.SendSmsInput{
		Txt: message,
		To:  &phone,
	}

	_, err := r.GetClient().SendSms(input)

	if err == nil {
		if r.logger != nil {
			r.logger.Info("Send success", xlog.F{
				"phone":   phone,
				"message": message,
			})
		}

		if r.metricTotalSendSuccess != nil {
			r.metricTotalSendSuccess.Add(1)
		}
	} else {
		if r.logger != nil {
			r.logger.Error("Send failed", xlog.F{
				"phone":   phone,
				"message": message,
				"error":   err.Error(),
			})
		}

		if r.metricTotalSendFailed != nil {
			r.metricTotalSendFailed.Add(1)
		}
	}

	return err
}

func (r *Resource) GetBalance() (float64, error) {
	info, err := r.GetClient().Info(nil)

	if err != nil {
		return -1, err
	}

	if r.metricBalance != nil {
		r.metricBalance.Set(info.Account)
	}

	return info.Account, nil
}
