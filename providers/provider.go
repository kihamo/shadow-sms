package providers

import (
	"context"
)

type Provider interface {
	Send(context.Context, string, string) (float64, error)
	Balance(context.Context) (float64, error)
}
