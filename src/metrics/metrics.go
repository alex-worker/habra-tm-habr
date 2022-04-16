package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func RunMetrics(profileAddress string) {
	log.Printf("Profile address: %v/debug/pprof/\n", profileAddress)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(profileAddress, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
