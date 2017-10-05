package main // import "github.com/kihamo/shadow-sms/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	sms "github.com/kihamo/shadow-sms/components/sms/instance"
	alerts "github.com/kihamo/shadow/components/alerts/instance"
	config "github.com/kihamo/shadow/components/config/instance"
	dashboard "github.com/kihamo/shadow/components/dashboard/instance"
	logger "github.com/kihamo/shadow/components/logger/instance"
	metrics "github.com/kihamo/shadow/components/metrics/instance"
)

func main() {
	application, err := shadow.NewApp(
		"Sms",
		"1.0",
		"12345-full",
		[]shadow.Component{
			sms.NewComponent(),
			alerts.NewComponent(),
			config.NewComponent(),
			dashboard.NewComponent(),
			logger.NewComponent(),
			metrics.NewComponent(),
		},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
