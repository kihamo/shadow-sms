package sms

import (
	"net/http"

	"github.com/kihamo/shadow/components/dashboard"
)

type IndexHandler struct {
	dashboard.Handler

	component *Component
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	balance, err := h.component.GetBalance()
	h.Render(r.Context(), "sms", "index", map[string]interface{}{
		"balanceError": err,
		"balanceValue": balance,
	})
}
