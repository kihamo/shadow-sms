package smsintel

import (
	"time"

	"github.com/kihamo/shadow/resource/config"
)

const (
	ConfigSmsLogin           = "sms.login"
	ConfigSmsPassword        = "sms.password"
	ConfigSmsMetricsInterval = "sms.metrics-interval"
)

func (r *Resource) GetConfigVariables() []config.Variable {
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

func (r *Resource) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigSmsLogin:           {r.watchLogin},
		ConfigSmsPassword:        {r.watchPassword},
		ConfigSmsMetricsInterval: {r.watchInterval},
	}
}

func (r *Resource) watchLogin(newValue interface{}, _ interface{}) {
	r.initClient(newValue.(string), r.config.GetString(ConfigSmsPassword))
}

func (r *Resource) watchPassword(newValue interface{}, _ interface{}) {
	r.initClient(r.config.GetString(ConfigSmsLogin), newValue.(string))
}

func (r *Resource) watchInterval(newValue interface{}, _ interface{}) {
	r.changeTicker <- newValue.(time.Duration)
}
