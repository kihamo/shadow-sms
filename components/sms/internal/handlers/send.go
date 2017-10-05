package handlers

import (
	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow/components/dashboard"
)

type SendHandler struct {
	dashboard.Handler

	Component sms.Component
}

func (h *SendHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if r.IsPost() {
		phone := r.Original().FormValue("phone")
		message := r.Original().FormValue("message")

		if err := h.Component.Send(message, phone); err != nil {
			w.SendJSON(map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			w.SendJSON(map[string]string{})
		}

		return
	}

	balance, err := h.Component.GetBalance()
	h.Render(r.Context(), h.Component.GetName(), "send", map[string]interface{}{
		"balanceError": err,
		"balanceValue": balance,
	})
}
