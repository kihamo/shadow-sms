package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/resource"
	r "github.com/kihamo/shadow/resource"
	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type SmsService struct {
	application *shadow.Application
	client      *smsintel.SmsIntel
	mutex       sync.RWMutex
	Info        *procedure.InfoOutput
}

func (s *SmsService) GetName() string {
	return "smsintel"
}

func (s *SmsService) Init(a *shadow.Application) error {
	s.application = a

	resourceSmsIntel, err := a.GetResource("smsintel")
	if err != nil {
		return err
	}

	s.client = resourceSmsIntel.(*resource.SmsIntel).GetClient()

	return nil
}

func (s *SmsService) Run() error {
	if s.application.HasResource("tasks") {
		tasks, _ := s.application.GetResource("tasks")
		tasks.(*r.Dispatcher).AddNamedTask("aws.updater", s.getInfoJob)
	}

	return nil
}

func (s *SmsService) getInfoJob(args ...interface{}) (bool, time.Duration) {
	var err error

	s.mutex.Lock()
	s.Info, err = s.client.Info(nil)

	fmt.Println(err, s.Info)

	s.mutex.Unlock()

	return true, time.Hour
}
