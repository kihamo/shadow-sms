package handlers

import (
	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow/components/dashboard"
)

type SendHandler struct {
	dashboard.Handler
}

func (h *SendHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	component := r.Component().(sms.Component)

	if r.IsPost() {
		phone := r.Original().FormValue("phone")
		message := r.Original().FormValue("message")

		if _, err := component.Send(message, phone); err != nil {
			w.SendJSON(map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			w.SendJSON(map[string]string{})
		}

		return
	}

	balance, err := component.GetBalance()
	h.Render(r.Context(), "send", map[string]interface{}{
		"balanceError": err,
		"balanceValue": balance,
	})
}
