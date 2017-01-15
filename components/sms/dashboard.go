package sms

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) GetDashboardMenu() *dashboard.Menu {
	return &dashboard.Menu{
		Name: "SMS",
		Url:  "/",
		Icon: "files-o",
		SubMenu: []*dashboard.Menu{{
			Name: "Balance",
			Url:  "/",
			Icon: "money",
		}, {
			Name: "Send",
			Url:  "/send",
			Icon: "send",
		}},
	}
}

func (c *Component) GetDashboardRoutes() []*dashboard.Route {
	return []*dashboard.Route{
		{
			Methods: []string{http.MethodGet, http.MethodPost},
			Path:    "/",
			Handler: &IndexHandler{
				component: c,
			},
		},
		{
			Methods: []string{http.MethodGet, http.MethodPost},
			Path:    "/send",
			Handler: &SendHandler{
				component: c,
			},
		},
	}
}
