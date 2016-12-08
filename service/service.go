package service

import (
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/resource/smsintel"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/workers"
	"github.com/rs/xlog"
)

type SmsService struct {
	application *shadow.Application
	sms         *smsintel.Resource
	logger      xlog.Logger

	mutex sync.RWMutex

	balanceValue float64
	balanceError error
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
	s.sms = resourceSmsIntel.(*smsintel.Resource)

	resourceLogger, err := a.GetResource("logger")
	if err != nil {
		return err
	}
	s.logger = resourceLogger.(*logger.Resource).Get(s.GetName())

	return nil
}

func (s *SmsService) Run() error {
	if s.application.HasResource("workers") {
		resourceWorkers, _ := s.application.GetResource("workers")
		resourceWorkers.(*workers.Resource).AddNamedTaskByFunc("sms.balance.updater", s.getBalanceJob)
	}

	return nil
}

func (s *SmsService) getBalanceJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, interface{}, error) {
	balance, err := s.sms.GetBalance()

	s.mutex.Lock()
	s.balanceError = err

	if err == nil {
		s.balanceValue = balance
	}
	s.mutex.Unlock()

	if err != nil && s.application.HasResource("alerts") {
		resourceAlerts, _ := s.application.GetResource("alerts")
		resourceAlerts.(*alerts.Resource).Send("Error get sms balance", err.Error(), "exclamation")

		return -1, time.Minute, nil, err
	}

	return -1, time.Hour, nil, nil
}
