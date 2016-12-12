package service

import (
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler

	service *SmsService
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("SMS")
	h.SetPageHeader("SMS")

	h.service.mutex.RLock()
	h.SetVar("BalanceError", h.service.balanceError)
	h.SetVar("BalanceValue", h.service.balanceValue)
	h.SetVar("BalancePositive", h.service.balanceValue > 0)
	h.service.mutex.RUnlock()
}
