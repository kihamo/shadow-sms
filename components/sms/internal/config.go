package internal

import (
	"time"

	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			sms.ConfigApiUrl,
			config.ValueTypeString,
			"http://lcab.smsintel.ru/lcabApi",
			"SMSIntel Api URL",
			true,
			nil,
			nil),
		config.NewVariable(
			sms.ConfigLogin,
			config.ValueTypeString,
			nil,
			"SMSIntel login",
			true,
			nil,
			nil),
		config.NewVariable(
			sms.ConfigPassword,
			config.ValueTypeString,
			nil,
			"SMSIntel password",
			true,
			nil,
			nil),
		config.NewVariable(
			sms.ConfigBalanceUpdateInterval,
			config.ValueTypeDuration,
			"1m",
			"Interval for balance updater",
			true,
			nil,
			nil),
		config.NewVariable(
			sms.ConfigInfoTimeout,
			config.ValueTypeDuration,
			"5s",
			"Timeout for info request",
			true,
			nil,
			nil),
		config.NewVariable(
			sms.ConfigSendTimeout,
			config.ValueTypeDuration,
			"5s",
			"Timeout for send request",
			true,
			nil,
			nil),
	}
}

func (c *Component) GetConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(c.GetName(), []string{sms.ConfigApiUrl, sms.ConfigLogin, sms.ConfigPassword}, c.watchForClient),
		config.NewWatcher(c.GetName(), []string{sms.ConfigBalanceUpdateInterval}, c.watchBalanceUpdateInterval),
	}
}

func (c *Component) watchForClient(_ string, newValue interface{}, _ interface{}) {
	c.initClient(c.config.GetString(sms.ConfigApiUrl), c.config.GetString(sms.ConfigLogin), c.config.GetString(sms.ConfigPassword))
}

func (c *Component) watchBalanceUpdateInterval(_ string, newValue interface{}, _ interface{}) {
	c.changeTicker <- newValue.(time.Duration)
}
