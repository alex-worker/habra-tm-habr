package proxy_app

import "habra-tm-habr/internal/pkg/network/http/metrics"

type AppConfig struct {
	Proxy    ProxyConfig
	Profile  metrics.ProfileConfig
	TextConf TextConfig
}

func GetDefaultConfig() AppConfig {
	proxyConf := GetDefaultProxyConfig()
	profileConf := metrics.GetDefaultConfig()
	textConf := GetDefaultTextConfig()

	return AppConfig{
		Proxy:    proxyConf,
		Profile:  profileConf,
		TextConf: textConf,
	}
}
