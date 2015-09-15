package service

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/service/frontend"
)

func (s *SmsService) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "templates",
	}
}

func (s *SmsService) GetFrontendMenu() *frontend.FrontendMenu {
	return &frontend.FrontendMenu{
		Name: "SMS",
		Url:  "/sms",
		Icon: "files-o",
	}
}

func (s *SmsService) SetFrontendHandlers(router *frontend.Router) {
	router.GET(s, "/sms", &IndexHandler{})
}
