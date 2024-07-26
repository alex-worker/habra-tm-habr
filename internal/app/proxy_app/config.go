package proxy_app

type ProxyConfig struct {
	ProxyAddress string
	SiteAddress  string
}

func GetDefaultProxyConfig() ProxyConfig {
	return ProxyConfig{
		ProxyAddress: ":8080",
		SiteAddress:  "http://habrahabr.ru",
	}
}
