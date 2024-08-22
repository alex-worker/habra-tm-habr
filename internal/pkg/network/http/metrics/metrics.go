package metrics

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

type ProfileConfig struct {
	ProfileAddress string
}

func GetDefaultConfig() ProfileConfig {
	return ProfileConfig{
		ProfileAddress: ":9090",
	}
}

func RunMetrics(c ProfileConfig) {
	log.Printf("Profile address: %v/debug/pprof/\n", c.ProfileAddress)
	err := http.ListenAndServe(c.ProfileAddress, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
