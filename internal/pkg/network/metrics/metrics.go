package metrics

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func RunMetrics(profileAddress string) {
	log.Printf("Profile address: %v/debug/pprof/\n", profileAddress)
	err := http.ListenAndServe(profileAddress, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
