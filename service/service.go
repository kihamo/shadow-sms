package service

import (
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
	if s.application.HasResource("workers") {
		workers, _ := s.application.GetResource("workers")
		workers.(*r.Workers).AddNamedTaskByFunc("sms.balance.updater", s.getBalanceJob)
	}

	return nil
}

func (s *SmsService) getBalanceJob(attempts int64, _ chan bool, args ...interface{}) (int64, time.Duration, error) {
	info, err := s.SmsClient.Info(nil)

	s.mutex.Lock()
	s.BalanceError = err

	if err == nil {
		s.BalanceValue = info.Account
	}
	s.mutex.Unlock()

	if err != nil {
		s.FrontendService.SendAlert("Error get sms balance", err.Error(), "exclamation")
		return -1, time.Minute, err
	}

	return -1, time.Hour, nil
}
