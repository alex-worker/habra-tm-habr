package proxy_app

type TextConfig struct {
	RunesInWorld int
}

func GetDefaultTextConfig() TextConfig {
	return TextConfig{
		RunesInWorld: 6,
	}
}
