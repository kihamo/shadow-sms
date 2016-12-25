package service

import (
	"github.com/kihamo/shadow-sms/resource/smsintel"
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler

	smsintel *smsintel.Resource
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.SetPageTitle("SMS")
	h.SetPageHeader("SMS")

	balance, err := h.smsintel.GetBalance()

	h.SetVar("BalanceError", err)
	h.SetVar("BalanceValue", balance)
	h.SetVar("BalancePositive", balance > 0)
}
