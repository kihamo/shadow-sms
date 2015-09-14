package service

import (
	"github.com/kihamo/shadow/service/frontend"
)

type IndexHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *IndexHandler) Handle() {
	h.SetTemplate("index.tpl.html")
	h.View.Context["PageTitle"] = "SMS"
	h.View.Context["PageHeader"] = "SMS"

	service := h.Service.(*SmsService)
	h.View.Context["Info"] = service.Info
	h.View.Context["BalancePositive"] = service.Info.Account > 0
}
