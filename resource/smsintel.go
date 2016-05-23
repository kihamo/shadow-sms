package resource

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type SmsIntel struct {
	application *shadow.Application
	config      *resource.Config
	client      *smsintel.SmsIntel
}

func (r *SmsIntel) GetName() string {
	return "smsintel"
}

func (r *SmsIntel) GetConfigVariables() []resource.ConfigVariable {
	return []resource.ConfigVariable{
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

func (r *SmsIntel) Init(a *shadow.Application) error {
	r.application = a
	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}

	r.config = resourceConfig.(*resource.Config)

	return nil
}

func (r *SmsIntel) GetClient() *smsintel.SmsIntel {
	if r.client == nil {
		r.client = smsintel.NewSmsIntel(r.config.GetString("sms.login"), r.config.GetString("sms.password"))
	}

	return r.client
}

func (r *SmsIntel) Send(message, phone string) error {
	input := &procedure.SendSmsInput{
		Txt: message,
		To:  &phone,
	}

	_, err := r.GetClient().SendSms(input)
	return err
}
