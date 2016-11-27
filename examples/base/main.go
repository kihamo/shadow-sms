package main // import "github.com/kihamo/shadow-sms/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	r "github.com/kihamo/shadow-sms/resource"
	s "github.com/kihamo/shadow-sms/service"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(resource.Config),
			new(resource.Logger),
			new(resource.Template),
			new(resource.Workers),
			new(alerts.Alert),
			new(r.SmsIntel),
		},
		[]shadow.Service{
			new(frontend.FrontendService),
			new(system.SystemService),
			new(s.SmsService),
		},
		"Sms",
		"1.0",
		"12345-full",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
