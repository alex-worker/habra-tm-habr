package proxy_app

type AppConfig struct {
	Proxy          ProxyConfig
	ProfileAddress string
	RunesInWorld   int
}

func GetDefaultConfig() AppConfig {
	proxyConf := GetDefaultProxyConfig()
	return AppConfig{
		Proxy:          proxyConf,
		RunesInWorld:   6,
		ProfileAddress: ":9090",
	}
}
