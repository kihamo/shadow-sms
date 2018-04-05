package internal

import (
	"time"

	"github.com/kihamo/shadow-sms/components/sms"
	"github.com/kihamo/shadow-sms/providers/terasms"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(sms.ConfigSmsIntelApiUrl, config.ValueTypeString).
			WithUsage("API URL").
			WithGroup("SMSIntel provider").
			WithEditable(true).
			WithDefault("http://lcab.smsintel.ru/lcabApi"),
		config.NewVariable(sms.ConfigSmsIntelLogin, config.ValueTypeString).
			WithUsage("Login").
			WithGroup("SMSIntel provider").
			WithEditable(true),
		config.NewVariable(sms.ConfigSmsIntelPassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("SMSIntel provider").
			WithEditable(true).
			WithView([]string{config.ViewPassword}),
		config.NewVariable(sms.ConfigTeraSmsApiUrl, config.ValueTypeString).
			WithUsage("API URL").
			WithGroup("TeraSms provider").
			WithEditable(true).
			WithDefault("https://auth.terasms.ru/"),
		config.NewVariable(sms.ConfigTeraSmsAuthType, config.ValueTypeInt).
			WithUsage("Auth type").
			WithGroup("TeraSms provider").
			WithEditable(true).
			WithDefault(terasms.AuthByToken).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{terasms.AuthByToken, "By token"},
					{terasms.AuthByLoginAndPassword, "By login and password"},
				},
			}),
		config.NewVariable(sms.ConfigTeraSmsLogin, config.ValueTypeString).
			WithUsage("Login").
			WithGroup("TeraSms provider").
			WithEditable(true),
		config.NewVariable(sms.ConfigTeraSmsPassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("TeraSms provider").
			WithEditable(true).
			WithView([]string{config.ViewPassword}),
		config.NewVariable(sms.ConfigTeraSmsToken, config.ValueTypeString).
			WithUsage("Token").
			WithGroup("TeraSms provider").
			WithEditable(true),
		config.NewVariable(sms.ConfigTeraSmsSender, config.ValueTypeString).
			WithUsage("Sender").
			WithGroup("TeraSms provider").
			WithEditable(true),
		config.NewVariable(sms.ConfigProvider, config.ValueTypeString).
			WithUsage("Provider").
			WithEditable(true).
			WithDefault(sms.ProviderSmsIntel).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{sms.ProviderSmsIntel, "SMSIntel"},
					{sms.ProviderTeraSms, "TeraSms"},
				},
			}),
		config.NewVariable(sms.ConfigBalanceUpdateInterval, config.ValueTypeDuration).
			WithUsage("Interval for balance updater").
			WithEditable(true).
			WithDefault("1m"),
		config.NewVariable(sms.ConfigTimeoutBalance, config.ValueTypeDuration).
			WithUsage("Timeout for info request").
			WithEditable(true).
			WithDefault("5s"),
		config.NewVariable(sms.ConfigTimeoutSend, config.ValueTypeDuration).
			WithUsage("Timeout for send request").
			WithEditable(true).
			WithDefault("5s"),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher([]string{sms.ConfigProvider}, c.watchProvider),
		config.NewWatcher([]string{
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
		config.NewWatcher([]string{sms.ConfigBalanceUpdateInterval}, c.watchBalanceUpdateInterval),
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
