package proxy_app

import "habra-tm-habr/internal/pkg/network/http/metrics"

type AppConfig struct {
	Proxy        ProxyConfig
	Profile      metrics.ProfileConfig
	RunesInWorld int
}

func GetDefaultConfig() AppConfig {
	proxyConf := GetDefaultProxyConfig()
	profileConf := metrics.GetDefaultConfig()

	return AppConfig{
		Proxy:        proxyConf,
		Profile:      profileConf,
		RunesInWorld: 6,
	}
}
