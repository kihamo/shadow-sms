package sms

import (
	"github.com/kihamo/shadow/components/dashboard"
)

type IndexHandler struct {
	dashboard.Handler

	component *Component
}

func (h *IndexHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if r.IsPost() {
		phone := r.Original().FormValue("phone")
		message := r.Original().FormValue("message")

		if err := h.component.Send(message, phone); err != nil {
			w.SendJSON(map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			w.SendJSON(map[string]string{})
		}

		return
	}

	balance, err := h.component.GetBalance()
	h.Render(r.Context(), ComponentName, "index", map[string]interface{}{
		"balanceError": err,
		"balanceValue": balance,
	})
}
