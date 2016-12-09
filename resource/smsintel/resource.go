package smsintel

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type Resource struct {
	application *shadow.Application
	config      *config.Resource
	client      *smsintel.SmsIntel
	logger      logger.Logger
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

	r.application = a

	return nil
}

func (r *Resource) Run() error {
	if resourceLogger, err := r.application.GetResource("logger"); err == nil {
		r.logger = resourceLogger.(*logger.Resource).Get(r.GetName())
	} else {
		r.logger = logger.NopLogger
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
		r.logger.Info("Send success", map[string]interface{}{
			"phone": phone,
			"text":  message,
		})

		if metricTotalSendSuccess != nil {
			metricTotalSendSuccess.Add(1)
		}
	} else {
		r.logger.Error("Send failed", map[string]interface{}{
			"phone": phone,
			"text":  message,
			"error": err.Error(),
		})

		if metricTotalSendFailed != nil {
			metricTotalSendFailed.Add(1)
		}
	}

	return err
}

func (r *Resource) GetBalance() (float64, error) {
	info, err := r.GetClient().Info(nil)

	if err != nil {
		return -1, err
	}

	if metricBalance != nil {
		metricBalance.Set(info.Account)
	}

	return info.Account, nil
}
