package smsintel

import (
	"github.com/Sirupsen/logrus"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type SmsIntel struct {
	application *shadow.Application
	client      *smsintel.SmsIntel
	config      *resource.Config
	logger      *logrus.Entry
	metrics     *metrics.Metrics
}

func (r *SmsIntel) GetName() string {
	return "smsintel"
}

func (r *SmsIntel) GetConfigVariables() []resource.ConfigVariable {
	return []resource.ConfigVariable{
		{
			Key:   "sms.login",
			Value: "",
			Usage: "SMSIntel login",
		},
		{
			Key:   "sms.password",
			Value: "",
			Usage: "SMSIntel password",
		},
	}
}

func (r *SmsIntel) Init(a *shadow.Application) error {
	r.application = a
	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}

	r.config = resourceConfig.(*resource.Config)

	resourceLogger, err := r.application.GetResource("logger")
	if err != nil {
		return err
	}
	r.logger = resourceLogger.(*resource.Logger).Get(r.GetName())

	if a.HasResource("metrics") {
		resourceMetrics, _ := a.GetResource("metrics")
		r.metrics = resourceMetrics.(*metrics.Metrics)
	}

	return nil
}

func (r *SmsIntel) GetClient() *smsintel.SmsIntel {
	if r.client == nil {
		r.client = smsintel.NewSmsIntel(r.config.GetString("sms.login"), r.config.GetString("sms.password"))
	}

	return r.client
}

func (r *SmsIntel) Send(message, phone string) error {
	input := &procedure.SendSmsInput{
		Txt: message,
		To:  &phone,
	}

	_, err := r.GetClient().SendSms(input)

	entry := r.logger.WithFields(logrus.Fields{
		"phone":   phone,
		"message": message,
	})

	if err == nil {
		entry.Info("Send success")

		if r.metrics != nil {
			r.metrics.NewCounter(MetricSmsTotalSendSuccess).Inc(1)
		}
	} else {
		entry.WithField("error", err.Error()).Error("Send failed")

		if r.metrics != nil {
			r.metrics.NewCounter(MetricSmsTotalSendFailed).Inc(1)
		}
	}

	return err
}

func (r *SmsIntel) GetBalance() (float64, error) {
	info, err := r.GetClient().Info(nil)

	if err != nil {
		return -1, err
	}

	if r.metrics != nil {
		r.metrics.NewGaugeFloat64(MetricSmsBalance).Update(info.Account)
	}

	return info.Account, nil
}
