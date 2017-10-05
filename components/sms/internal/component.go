package internal

import (
	"context"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow/components/alerts"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type Component struct {
	application shadow.Application
	alerts      alerts.Component
	config      config.Component
	logger      logger.Logger
	routes      []dashboard.Route

	mutex        sync.RWMutex
	client       *smsintel.SmsIntel
	changeTicker chan time.Duration
}

func (c *Component) GetName() string {
	return sms.ComponentName
}

func (c *Component) GetVersion() string {
	return sms.ComponentVersion
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
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	c.application = a
	c.changeTicker = make(chan time.Duration)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	c.initClient(c.config.GetString(sms.ConfigApiUrl), c.config.GetString(sms.ConfigLogin), c.config.GetString(sms.ConfigPassword))

	if cmpAlerts := c.application.GetComponent(alerts.ComponentName); cmpAlerts != nil {
		c.alerts = cmpAlerts.(alerts.Component)
	}

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(c.config.GetDuration(sms.ConfigBalanceUpdateInterval))

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

func (c *Component) initClient(apiUrl, login, password string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.client = smsintel.NewSmsIntel(login, password)
	c.client.SetOptions(map[string]string{
		"api_url": apiUrl,
	})
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

	var err error

	timeout := c.config.GetDuration(sms.ConfigSendTimeout)
	if timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		_, err = c.GetClient().SendSmsWithContext(ctx, input)
	} else {
		_, err = c.GetClient().SendSms(input)
	}

	if err == nil {
		c.logger.Info("Send success", map[string]interface{}{
			"phone": phone,
			"text":  message,
		})

		if metricTotalSend != nil {
			metricTotalSend.With("status", "success").Inc()
		}
	} else {
		c.logger.Error("Send failed", map[string]interface{}{
			"phone": phone,
			"text":  message,
			"error": err.Error(),
		})

		if metricTotalSend != nil {
			metricTotalSend.With("status", "failed").Inc()
		}
	}

	return err
}

func (c *Component) GetBalance() (float64, error) {
	var (
		info *procedure.InfoOutput
		err  error
	)

	timeout := c.config.GetDuration(sms.ConfigInfoTimeout)
	if timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		info, err = c.GetClient().InfoWithContext(ctx, nil)
	} else {
		info, err = c.GetClient().Info(nil)
	}

	if err != nil {
		return -1, err
	}

	return info.Account, nil
}
