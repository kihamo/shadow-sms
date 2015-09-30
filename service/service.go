package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/resource"
	r "github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/smsintel"
)

type SmsService struct {
	application *shadow.Application
	mutex       sync.RWMutex

	FrontendService *frontend.FrontendService
	SmsClient       *smsintel.SmsIntel
	Logger          *logrus.Entry
	BalanceValue    float64
	BalanceError    error
}

func (s *SmsService) GetName() string {
	return "sms"
}

func (s *SmsService) Init(a *shadow.Application) error {
	s.application = a

	serviceFrontend, err := a.GetService("frontend")
	if err != nil {
		return err
	}
	s.FrontendService = serviceFrontend.(*frontend.FrontendService)

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

func (s *SmsService) Run() error {
	if s.application.HasResource("tasks") {
		tasks, _ := s.application.GetResource("tasks")
		tasks.(*r.Dispatcher).AddNamedTask("sms.balance.updater", s.getBalanceJob)
	}

	return nil
}

func (s *SmsService) getBalanceJob(args ...interface{}) (repeat int64, duration time.Duration) {
	info, err := s.SmsClient.Info(nil)

	s.mutex.Lock()
	s.BalanceError = err

	if err == nil {
		s.BalanceValue = info.Account
	}
	s.mutex.Unlock()

	if s.BalanceError != nil {
		s.Logger.Warn(s.BalanceError.Error())
		s.FrontendService.SendAlert("Error get sms balance", s.BalanceError.Error(), "exclamation")

		duration = time.Minute
	} else {
		duration = time.Hour
	}

	return -1, duration
}
