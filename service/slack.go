package service

import (
	slacks "github.com/kihamo/shadow-slack/service"
)

func (s *SmsService) GetSlackCommands() []slacks.SlackCommand {
	return []slacks.SlackCommand{
		&BalanceCommand{},
	}
}
