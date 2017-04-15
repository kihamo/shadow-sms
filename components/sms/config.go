package sms

import (
	"time"

	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigLogin                 = ComponentName + ".login"
	ConfigPassword              = ComponentName + ".password"
	ConfigBalanceUpdateInterval = ComponentName + ".balance-updater-interval"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:      ConfigLogin,
			Usage:    "SMSIntel login",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigPassword,
			Usage:    "SMSIntel password",
			Type:     config.ValueTypeString,
			Editable: true,
		},
		{
			Key:      ConfigBalanceUpdateInterval,
			Usage:    "Interval for balance updater",
			Default:  "1m",
			Type:     config.ValueTypeDuration,
			Editable: true,
		},
	}
}

func (c *Component) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigLogin:                 {c.watchLogin},
		ConfigPassword:              {c.watchPassword},
		ConfigBalanceUpdateInterval: {c.watchBalanceUpdateInterval},
	}
}

func (c *Component) watchLogin(_ string, newValue interface{}, _ interface{}) {
	c.initClient(newValue.(string), c.config.GetString(ConfigPassword))
}

func (c *Component) watchPassword(_ string, newValue interface{}, _ interface{}) {
	c.initClient(c.config.GetString(ConfigLogin), newValue.(string))
}

func (c *Component) watchBalanceUpdateInterval(_ string, newValue interface{}, _ interface{}) {
	c.changeTicker <- newValue.(time.Duration)
}
