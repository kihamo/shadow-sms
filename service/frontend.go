package service

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/service/frontend"
)

func (s *SmsService) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (s *SmsService) GetFrontendMenu() *frontend.FrontendMenu {
	return &frontend.FrontendMenu{
		Name: "SMS",
		Url:  "/sms",
		Icon: "files-o",
		SubMenu: []*frontend.FrontendMenu{{
			Name: "Balance",
			Url:  "/sms",
			Icon: "send",
		}, {
			Name: "Send",
			Url:  "/sms/send",
			Icon: "send",
		}},
	}
}

func (s *SmsService) SetFrontendHandlers(router *frontend.Router) {
	router.GET(s, "/sms", &IndexHandler{
		smsintel: s.sms,
	})

	handlerSend := &SendHandler{
		smsintel: s.sms,
	}
	router.GET(s, "/sms/send", handlerSend)
	router.POST(s, "/sms/send", handlerSend)
}
