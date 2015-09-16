package service

import (
	"github.com/kihamo/shadow"
	slacks "github.com/kihamo/shadow-slack/service"
	sl "github.com/nlopes/slack"
	"gopkg.in/jcelliott/turnpike.v2"
)

type BalanceCommand struct {
	slacks.AbstractSlackCommand
	Service *SmsService

	client *turnpike.Client
}

func (c *BalanceCommand) GetName() string {
	return "sms.balance"
}

func (c *BalanceCommand) GetDescription() string {
	return "Get balance in SMSIntel account"
}

func (c *BalanceCommand) Init(s shadow.Service, a *shadow.Application) {
	c.AbstractSlackCommand.Init(s, a)
	c.Service = s.(*SmsService)
}

func (c *BalanceCommand) Run(m *sl.MessageEvent, args ...string) {
	if c.Service.BalanceError == nil {
		c.SendMessagef(m.Channel, "%.2f rub", c.Service.BalanceValue)
	} else {
		c.SendMessage(m.Channel, "Unknown balance")
	}
}
