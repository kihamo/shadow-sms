package sms

import (
	"github.com/kihamo/shadow"
)

type Component interface {
	shadow.Component

	Send(message, phone string) (float64, error)
	GetBalance() (float64, error)
}
