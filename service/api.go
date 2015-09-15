package service

import (
	apis "github.com/kihamo/shadow-api/service"
)

func (s *SmsService) GetApiProcedures() []apis.ApiProcedure {
	return []apis.ApiProcedure{
		&SendProcedure{},
	}
}
