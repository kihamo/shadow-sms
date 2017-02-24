package sms

import (
	"time"

	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigSmsLogin           = "sms.login"
	ConfigSmsPassword        = "sms.password"
	ConfigSmsMetricsInterval = "sms.metrics-interval"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:      ConfigSmsLogin,
			Usage:    "SMSIntel login",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigSmsPassword,
			Usage:    "SMSIntel password",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigSmsMetricsInterval,
			Usage:    "Interval for balance updater",
			Default:  "1m",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigSmsLogin:           {c.watchLogin},
		ConfigSmsPassword:        {c.watchPassword},
		ConfigSmsMetricsInterval: {c.watchInterval},
	}
}

func (c *Component) watchLogin(_ string, newValue interface{}, _ interface{}) {
	c.initClient(newValue.(string), c.config.GetString(ConfigSmsPassword))
}

func (c *Component) watchPassword(_ string, newValue interface{}, _ interface{}) {
	c.initClient(c.config.GetString(ConfigSmsLogin), newValue.(string))
}

func (c *Component) watchInterval(_ string, newValue interface{}, _ interface{}) {
	c.changeTicker <- newValue.(time.Duration)
}
