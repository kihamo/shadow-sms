package service

import (
	"github.com/kihamo/shadow/service/api"
)

func (s *SmsService) GetApiProcedures() []api.ApiProcedure {
	return []api.ApiProcedure{
		&SendProcedure{},
	}
}
