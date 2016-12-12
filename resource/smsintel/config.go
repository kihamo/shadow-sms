package smsintel

import (
	"github.com/kihamo/shadow/resource/config"
)

const (
	ConfigSmsLogin    = "sms.login"
	ConfigSmsPassword = "sms.password"
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
	}
}

func (r *Resource) GetConfigWatchers() map[string][]config.Watcher {
	return map[string][]config.Watcher{
		ConfigSmsLogin:    {r.watchAuth},
		ConfigSmsPassword: {r.watchAuth},
	}
}

func (r *Resource) watchAuth(_ interface{}, _ interface{}) {
	r.initClient()
}
