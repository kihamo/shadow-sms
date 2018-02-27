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
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
)

type Component struct {
	application shadow.Application
	config      config.Component
	logger      logger.Logger
	routes      []dashboard.Route

	mutex         sync.RWMutex
	provider      providers.Provider
	changeTicker  chan time.Duration
	metricEnabled bool
}

func (c *Component) Name() string {
	return sms.ComponentName
}

func (c *Component) Version() string {
	return sms.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	c.application = a
	c.changeTicker = make(chan time.Duration)
	c.metricEnabled = a.HasComponent(metrics.ComponentName)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.Name(), c.application)

	c.initProvider()

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(c.config.Duration(sms.ConfigBalanceUpdateInterval))

		for {
			select {
			case <-ticker.C:
				balance, err := c.GetBalance()

				if err != nil {
					c.logger.Error("Get SMS balance failed", map[string]interface{}{
						"error": err.Error(),
					})
				} else if c.metricEnabled {
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

	id := c.config.String(sms.ConfigProvider)

	switch id {
	case sms.ProviderSmsIntel:
		p, err = smsintel.NewClient(
			c.config.String(sms.ConfigSmsIntelApiUrl),
			c.config.String(sms.ConfigSmsIntelLogin),
			c.config.String(sms.ConfigSmsIntelPassword))

	case sms.ProviderTeraSms:
		p, err = terasms.NewClient(
			c.config.String(sms.ConfigTeraSmsApiUrl),
			c.config.Int(sms.ConfigTeraSmsAuthType),
			c.config.String(sms.ConfigTeraSmsLogin),
			c.config.String(sms.ConfigTeraSmsPassword),
			c.config.String(sms.ConfigTeraSmsToken),
			c.config.String(sms.ConfigTeraSmsSender))
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

	timeout := c.config.Duration(sms.ConfigTimeoutSend)
	if timeout > 0 {
		ctx, ctxCancel = context.WithTimeout(ctx, timeout)
	}
	defer ctxCancel()

	err := c.GetProvider().Send(ctx, phone, message)
	if err == nil {
		c.logger.Debug("Send success", map[string]interface{}{
			"phone": phone,
			"text":  message,
		})

		if c.metricEnabled {
			metricTotalSend.With("status", "success").Inc()
		}
	} else {
		c.logger.Error("Send failed", map[string]interface{}{
			"phone": phone,
			"text":  message,
			"error": err.Error(),
		})

		if c.metricEnabled {
			metricTotalSend.With("status", "failed").Inc()
		}
	}

	return err
}

func (c *Component) GetBalance() (float64, error) {
	ctx := context.Background()
	var ctxCancel func()

	timeout := c.config.Duration(sms.ConfigTimeoutBalance)
	if timeout > 0 {
		ctx, ctxCancel = context.WithTimeout(ctx, timeout)
	}
	defer ctxCancel()

	return c.GetProvider().Balance(ctx)
}
