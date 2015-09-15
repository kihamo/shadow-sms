package service

import (
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler

	balanceValue float64
	balanceError error

	mutex sync.RWMutex
}

func (h *IndexHandler) Init(a *shadow.Application, s shadow.Service) {
	if a.HasResource("tasks") {
		tasks, _ := a.GetResource("tasks")
		tasks.(*resource.Dispatcher).AddNamedTask("sms.balance.updater", h.getInfoJob)
	}

	h.AbstractFrontendHandler.Init(a, s)
}

func (h *IndexHandler) getInfoJob(args ...interface{}) (int64, time.Duration) {
	service := h.Service.(*SmsService)
	info, err := service.SmsClient.Info(nil)

	h.mutex.Lock()
	h.balanceError = err

	if err == nil {
		h.balanceValue = info.Account
	}
	h.mutex.Unlock()

	if h.balanceError != nil {
		service.Logger.Warn(h.balanceError.Error())
	}

	return -1, time.Hour
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.View.Context["PageTitle"] = "SMS"
	h.View.Context["PageHeader"] = "SMS"

	h.View.Context["BalanceValue"] = h.balanceValue
	h.View.Context["BalanceError"] = h.balanceError
	h.View.Context["BalancePositive"] = h.balanceValue > 0
}
