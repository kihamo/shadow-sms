package service

import (
	"github.com/kihamo/shadow-sms/resource/smsintel"
	"github.com/kihamo/shadow/service/frontend"
)

type SendHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *SendHandler) Handle() {
	if h.IsPost() {
		phone := h.Input.FormValue("phone")
		message := h.Input.FormValue("message")

		resourceSms, _ := h.Application.GetResource("smsintel")
		if err := resourceSms.(*smsintel.Resource).Send(message, phone); err != nil {
			h.SendJSON(map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			h.SendJSON(map[string]string{})
		}

		return
	}

	h.SetTemplate("send.tpl.html")
	h.SetPageTitle("Send sms")
	h.SetPageHeader("Send sms")
}
