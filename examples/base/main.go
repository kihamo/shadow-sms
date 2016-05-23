package main // import "github.com/kihamo/shadow-sms/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	smsr "github.com/kihamo/shadow-sms/resource"
	smss "github.com/kihamo/shadow-sms/service"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/service/frontend"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(resource.Config),
			new(resource.Logger),
			new(resource.Template),
			new(resource.Dispatcher),
			new(smsr.SmsIntel),
		},
		[]shadow.Service{
			new(frontend.FrontendService),
			new(smss.SmsService),
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
