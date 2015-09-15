package service

import (
	"github.com/Sirupsen/logrus"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/resource"
	r "github.com/kihamo/shadow/resource"
	"github.com/kihamo/smsintel"
)

type SmsService struct {
	application *shadow.Application

	SmsClient *smsintel.SmsIntel
	Logger    *logrus.Entry
}

func (s *SmsService) GetName() string {
	return "sms"
}

func (s *SmsService) Init(a *shadow.Application) error {
	s.application = a

	resourceSmsIntel, err := a.GetResource("smsintel")
	if err != nil {
		return err
	}

	s.SmsClient = resourceSmsIntel.(*resource.SmsIntel).GetClient()

	resourceLogger, err := a.GetResource("logger")
	if err != nil {
		return err
	}
	logger := resourceLogger.(*r.Logger)
	s.Logger = logger.Get(s.GetName())

	return nil
}
