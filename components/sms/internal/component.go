package internal

import (
	"context"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow-sms/providers"
	"github.com/kihamo/shadow-sms/providers/smsintel"
	"github.com/kihamo/shadow-sms/providers/terasms"
	"github.com/kihamo/shadow/components/alerts"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
)

type Component struct {
	application shadow.Application
	alerts      alerts.Component
	config      config.Component
	logger      logger.Logger
	routes      []dashboard.Route

	mutex        sync.RWMutex
	provider     providers.Provider
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

	c.initProvider()

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

func (c *Component) initProvider() {
	var (
		p   providers.Provider
		err error
	)

	id := c.config.GetString(sms.ConfigProvider)

	switch id {
	case sms.ProviderSmsIntel:
		p, err = smsintel.NewClient(
			c.config.GetString(sms.ConfigSmsIntelApiUrl),
			c.config.GetString(sms.ConfigSmsIntelLogin),
			c.config.GetString(sms.ConfigSmsIntelPassword))

	case sms.ProviderTeraSms:
		p, err = terasms.NewClient(
			c.config.GetString(sms.ConfigTeraSmsApiUrl),
			c.config.GetInt(sms.ConfigTeraSmsAuthType),
			c.config.GetString(sms.ConfigTeraSmsLogin),
			c.config.GetString(sms.ConfigTeraSmsPassword),
			c.config.GetString(sms.ConfigTeraSmsToken),
			c.config.GetString(sms.ConfigTeraSmsSender))
	}

	if err == nil {
		c.mutex.Lock()
		c.provider = p
		c.mutex.Unlock()
	} else {
		c.logger.Errorf("Failed init sms provider %s with error %s", id, err.Error())
	}
}

func (c *Component) GetProvider() providers.Provider {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.provider
}

func (c *Component) Send(message, phone string) error {
	ctx := context.Background()
	var ctxCancel func()

	timeout := c.config.GetDuration(sms.ConfigTimeoutSend)
	if timeout > 0 {
		ctx, ctxCancel = context.WithTimeout(ctx, timeout)
	}
	defer ctxCancel()

	err := c.GetProvider().Send(ctx, phone, message)
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
	ctx := context.Background()
	var ctxCancel func()

	timeout := c.config.GetDuration(sms.ConfigTimeoutBalance)
	if timeout > 0 {
		ctx, ctxCancel = context.WithTimeout(ctx, timeout)
	}
	defer ctxCancel()

	return c.GetProvider().Balance(ctx)
}
