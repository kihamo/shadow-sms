package sms

import (
	"net/http"

	"github.com/kihamo/shadow/components/dashboard"
)

type SendHandler struct {
	dashboard.Handler

	component *Component
}

func (h *SendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.IsPost(r) {
		phone := r.FormValue("phone")
		message := r.FormValue("message")

		if err := h.component.Send(message, phone); err != nil {
			h.SendJSON(map[string]interface{}{
				"error": err.Error(),
			}, w)
		} else {
			h.SendJSON(map[string]string{}, w)
		}

		return
	}

	h.Render(r.Context(), "sms", "send", nil)
}
