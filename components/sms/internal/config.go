package internal

import (
	"time"

	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow-sms/providers/terasms"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) GetConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			sms.ConfigProvider,
			config.ValueTypeString,
			sms.ProviderSmsIntel,
			"Sms Provider",
			true,
			"Others",
			[]string{
				config.ViewEnum,
			},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{sms.ProviderSmsIntel, "SMSIntel"},
					{sms.ProviderTeraSms, "TeraSms"},
				},
			}),
		config.NewVariable(
			sms.ConfigSmsIntelApiUrl,
			config.ValueTypeString,
			"http://lcab.smsintel.ru/lcabApi",
			"SMSIntel Api URL",
			true,
			"SMSIntel provider",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigSmsIntelLogin,
			config.ValueTypeString,
			nil,
			"SMSIntel login",
			true,
			"SMSIntel provider",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigSmsIntelPassword,
			config.ValueTypeString,
			nil,
			"SMSIntel password",
			true,
			"SMSIntel provider",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			sms.ConfigTeraSmsApiUrl,
			config.ValueTypeString,
			"https://auth.terasms.ru/",
			"TeraSms Api URL",
			true,
			"TeraSms provider",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigTeraSmsAuthType,
			config.ValueTypeInt,
			terasms.AuthByToken,
			"Sms auth type",
			true,
			"TeraSms provider",
			[]string{
				config.ViewEnum,
			},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{terasms.AuthByToken, "By token"},
					{terasms.AuthByLoginAndPassword, "By login and password"},
				},
			}),
		config.NewVariable(
			sms.ConfigTeraSmsLogin,
			config.ValueTypeString,
			nil,
			"TeraSms login",
			true,
			"TeraSms provider",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigTeraSmsPassword,
			config.ValueTypeString,
			nil,
			"TeraSms password",
			true,
			"TeraSms provider",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			sms.ConfigTeraSmsToken,
			config.ValueTypeString,
			nil,
			"TeraSms token",
			true,
			"TeraSms provider",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			sms.ConfigTeraSmsSender,
			config.ValueTypeString,
			nil,
			"TeraSms sender",
			true,
			"TeraSms provider",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigBalanceUpdateInterval,
			config.ValueTypeDuration,
			"1m",
			"Interval for balance updater",
			true,
			"Others",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigTimeoutBalance,
			config.ValueTypeDuration,
			"5s",
			"Timeout for info request",
			true,
			"Others",
			nil,
			nil),
		config.NewVariable(
			sms.ConfigTimeoutSend,
			config.ValueTypeDuration,
			"5s",
			"Timeout for send request",
			true,
			"Others",
			nil,
			nil),
	}
}

func (c *Component) GetConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(c.GetName(), []string{sms.ConfigProvider}, c.watchProvider),
		config.NewWatcher(c.GetName(), []string{
			sms.ConfigSmsIntelApiUrl,
			sms.ConfigSmsIntelLogin,
			sms.ConfigSmsIntelPassword,
			sms.ConfigTeraSmsApiUrl,
			sms.ConfigTeraSmsAuthType,
			sms.ConfigTeraSmsLogin,
			sms.ConfigTeraSmsPassword,
			sms.ConfigTeraSmsToken,
			sms.ConfigTeraSmsSender},
			c.watchReinitProvider),
		config.NewWatcher(c.GetName(), []string{sms.ConfigBalanceUpdateInterval}, c.watchBalanceUpdateInterval),
	}
}

func (c *Component) watchProvider(_ string, newValue interface{}, oldValue interface{}) {
	if newValue != oldValue {
		c.logger.Infof("SMS provider changed from %s to %s", oldValue, newValue)
	}

	c.initProvider()
}

func (c *Component) watchReinitProvider(_ string, newValue interface{}, _ interface{}) {
	c.initProvider()
}

func (c *Component) watchBalanceUpdateInterval(_ string, newValue interface{}, _ interface{}) {
	c.changeTicker <- newValue.(time.Duration)
}
