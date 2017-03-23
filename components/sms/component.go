package sms

import (
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/alerts"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

const (
	ComponentName = "sms"
)

type Component struct {
	application shadow.Application
	alerts      *alerts.Component
	config      *config.Component
	logger      logger.Logger

	mutex        sync.RWMutex
	client       *smsintel.SmsIntel
	changeTicker chan time.Duration
}

func (c *Component) GetName() string {
	return ComponentName
}

func (c *Component) GetVersion() string {
	return ComponentVersion
}

func (c *Component) GetDependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name: alerts.ComponentName,
		},
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: logger.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.config = a.GetComponent(config.ComponentName).(*config.Component)

	c.application = a
	c.changeTicker = make(chan time.Duration)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	c.initClient(c.config.GetString(ConfigSmsLogin), c.config.GetString(ConfigSmsPassword))

	if cmpAlerts := c.application.GetComponent(alerts.ComponentName); cmpAlerts != nil {
		c.alerts = cmpAlerts.(*alerts.Component)
	}

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(c.config.GetDuration(ConfigSmsMetricsInterval))

		for {
			select {
			case <-ticker.C:
				balance, err := c.GetBalance()

				if err != nil {
					if c.alerts != nil {
						c.alerts.Send("Error get sms balance", err.Error(), "exclamation")
					}
				} else if metricBalance != nil {
					metricBalance.Set(balance)
				}

			case d := <-c.changeTicker:
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

func (c *Component) initClient(login, password string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.client = smsintel.NewSmsIntel(login, password)
}

func (c *Component) GetClient() *smsintel.SmsIntel {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.client
}

func (c *Component) Send(message, phone string) error {
	input := &procedure.SendSmsInput{
		Txt: message,
		To:  &phone,
	}

	_, err := c.GetClient().SendSms(input)

	if err == nil {
		c.logger.Info("Send success", map[string]interface{}{
			"phone": phone,
			"text":  message,
		})

		if metricTotalSendSuccess != nil {
			metricTotalSendSuccess.Add(1)
		}
	} else {
		c.logger.Error("Send failed", map[string]interface{}{
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

func (c *Component) GetBalance() (float64, error) {
	info, err := c.GetClient().Info(nil)

	if err != nil {
		return -1, err
	}

	return info.Account, nil
}
