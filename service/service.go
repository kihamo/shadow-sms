package service

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/resource/smsintel"
)

type SmsService struct {
	application *shadow.Application
	sms         *smsintel.Resource
}

func (s *SmsService) GetName() string {
	return "sms"
}

func (s *SmsService) Init(a *shadow.Application) error {
	resourceSmsIntel, err := a.GetResource("smsintel")
	if err != nil {
		return err
	}
	s.sms = resourceSmsIntel.(*smsintel.Resource)

	s.application = a

	return nil
}
