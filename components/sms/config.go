package sms

import (
	"time"

	"github.com/kihamo/shadow/components/config"
)

const (
	ConfigApiUrl                = ComponentName + ".api-url"
	ConfigLogin                 = ComponentName + ".login"
	ConfigPassword              = ComponentName + ".password"
	ConfigBalanceUpdateInterval = ComponentName + ".balance-updater-interval"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:      ConfigApiUrl,
			Usage:    "SMSIntel Api URL",
			Default:  "http://lcab.smsintel.ru/lcabApi",
			Type:     config.ValueTypeString,
			Editable: true,
		},
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
		ConfigApiUrl:                {c.watchApiUrl},
		ConfigLogin:                 {c.watchLogin},
		ConfigPassword:              {c.watchPassword},
		ConfigBalanceUpdateInterval: {c.watchBalanceUpdateInterval},
	}
}

func (c *Component) watchApiUrl(_ string, newValue interface{}, _ interface{}) {
	c.initClient(newValue.(string), c.config.GetString(ConfigLogin), c.config.GetString(ConfigPassword))
}

func (c *Component) watchLogin(_ string, newValue interface{}, _ interface{}) {
	c.initClient(c.config.GetString(ConfigApiUrl), newValue.(string), c.config.GetString(ConfigPassword))
}

func (c *Component) watchPassword(_ string, newValue interface{}, _ interface{}) {
	c.initClient(c.config.GetString(ConfigApiUrl), c.config.GetString(ConfigLogin), newValue.(string))
}

func (c *Component) watchBalanceUpdateInterval(_ string, newValue interface{}, _ interface{}) {
	c.changeTicker <- newValue.(time.Duration)
}
