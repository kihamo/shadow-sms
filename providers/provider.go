package providers

import (
	"context"
)

type Provider interface {
	Send(context.Context, string, string) error
	Balance(context.Context) (float64, error)
}
