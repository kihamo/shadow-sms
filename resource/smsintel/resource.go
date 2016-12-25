package smsintel

import (
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type Resource struct {
	application *shadow.Application

	alerts *alerts.Resource
	config *config.Resource
	logger logger.Logger

	mutex        sync.RWMutex
	client       *smsintel.SmsIntel
	changeTicker chan time.Duration
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

	r.changeTicker = make(chan time.Duration)

	return nil
}

func (r *Resource) Run(wg *sync.WaitGroup) error {
	r.logger = logger.NewOrNop(r.GetName(), r.application)

	r.initClient(r.config.GetString(ConfigSmsLogin), r.config.GetString(ConfigSmsPassword))

	resourceAlerts, err := r.application.GetResource("alerts")
	if err == nil {
		r.alerts = resourceAlerts.(*alerts.Resource)
	}

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(r.config.GetDuration(ConfigSmsMetricsInterval))

		for {
			select {
			case <-ticker.C:
				balance, err := r.GetBalance()

				if err != nil {
					if r.alerts != nil {
						r.alerts.Send("Error get sms balance", err.Error(), "exclamation")
					}
				} else if metricBalance != nil {
					metricBalance.Set(balance)
				}

			case d := <-r.changeTicker:
				if d.Nanoseconds() > 0 {
					ticker = time.NewTicker(d)
				} else {
					ticker.Stop()
				}
			}
		}
	}()

	return nil
}

func (r *Resource) initClient(login, password string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.client = smsintel.NewSmsIntel(login, password)
}

func (r *Resource) GetClient() *smsintel.SmsIntel {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

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

	return info.Account, nil
}
