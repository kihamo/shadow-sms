package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/service/api"
	"github.com/kihamo/smsintel/procedure"
	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	SendAttempts = 5
)

type SendRequest struct {
	Message string   `json:"message"`
	Phones  []string `json:"phones"`
}

type SendProcedure struct {
	api.AbstractApiProcedure

	tasks *resource.Dispatcher
}

func (p *SendProcedure) Init(s shadow.Service, a *shadow.Application) {
	p.AbstractApiProcedure.Init(s, a)

	if a.HasResource("tasks") {
		resourceTasks, _ := a.GetResource("tasks")
		p.tasks = resourceTasks.(*resource.Dispatcher)
	}
}

func (p *SendProcedure) GetName() string {
	return "sms.send"
}

func (p *SendProcedure) GetRequest() interface{} {
	return &SendRequest{}
}

func (p *SendProcedure) Run(r interface{}) *turnpike.CallResult {
	request := r.(*SendRequest)
	client := p.Service.(*SmsService).SmsClient

	p.tasks.AddNamedTask(p.GetName(), p.jobSend, request.Message, request.Phones)

	fmt.Println(request, client)

	return p.GetResult(nil, map[string]interface{}{
		"result": "success",
	})
}

func (p *SendProcedure) jobSend(args ...interface{}) (repeat int64, period time.Duration) {
	message := args[0].(string)
	phones := args[1].([]string)

	_, err := p.Service.(*SmsService).SmsClient.SendSms(&procedure.SendSmsInput{
		Txt: message,
		To:  strings.Join(phones, ","),
	})

	if err != nil {
		repeat = SendAttempts
	}

	return repeat, time.Minute * 5
}
