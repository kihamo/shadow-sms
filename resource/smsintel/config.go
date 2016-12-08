package smsintel

import (
	"github.com/kihamo/shadow/resource/config"
)

func (r *Resource) GetConfigVariables() []config.Variable {
	return []config.Variable{
		{
			Key:   "sms.login",
			Value: "",
			Usage: "SMSIntel login",
		},
		{
			Key:   "sms.password",
			Value: "",
			Usage: "SMSIntel password",
		},
	}
}
