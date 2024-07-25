package config

type Config struct {
	ProfileAddress string
	ProxyAddress   string
	SiteAddress    string
	RunesInWorld   int
	MetricsEnabled bool
}

func New() Config {
	return Config{
		ProfileAddress: ":9090",
		ProxyAddress:   ":8080",
		SiteAddress:    "http://habrahabr.ru",
		RunesInWorld:   6,
		MetricsEnabled: true,
	}
}
