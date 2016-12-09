package service

import (
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("SMS")
	h.SetPageHeader("SMS")

	service := h.Service.(*SmsService)

	service.mutex.RLock()
	h.SetVar("BalanceError", service.balanceError)
	h.SetVar("BalanceValue", service.balanceValue)
	h.SetVar("BalancePositive", service.balanceValue > 0)
	service.mutex.RUnlock()
}
