package app

import (
	"habra-tm-habr/config"
	"habra-tm-habr/internal/pkg/network/http/http_handler"
	"habra-tm-habr/internal/pkg/network/http/request_handler"
	"habra-tm-habr/internal/pkg/network/http/response_handler"
	"habra-tm-habr/internal/pkg/network/metrics"
	"habra-tm-habr/internal/pkg/text_processor"
	"log"
	"net/http"
)

type App struct {
	conf config.Config
}

type InterfaceApp interface {
	Run()
}

func New(c config.Config) InterfaceApp {
	app := &App{
		c,
	}
	return app
}

func (a *App) Run() {
	log.Printf("Application run...")

	if a.conf.MetricsEnabled {
		go metrics.RunMetrics(a.conf.ProfileAddress)
	}
	log.Printf("Proxy address %v -> %v\n", a.conf.ProxyAddress, a.conf.SiteAddress)

	handlerRequest := request_handler.NewRequestProxyHandler(a.conf.SiteAddress)

	handlerRaw := response_handler.NewHandlerRaw()

	p := text_processor.NewTmProcessor(a.conf.RunesInWorld)
	handlerHttp := response_handler.NewHandlerHtml(p)

	myHandler := http_handler.NewProxyHandler(handlerRequest, handlerRaw, handlerHttp)

	err := http.ListenAndServe(a.conf.ProxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
