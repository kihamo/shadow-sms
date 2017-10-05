package instance

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-sms/components/sms/internal"
)

func NewComponent() shadow.Component {
	return &internal.Component{}
}
