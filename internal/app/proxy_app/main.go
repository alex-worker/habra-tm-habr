package proxy_app

import (
	"habra-tm-habr/internal/app"
	UseCase "habra-tm-habr/internal/pkg"
	"habra-tm-habr/internal/pkg/network/http/metrics"
	"log"
)

func New(c AppConfig) app.ApplicationInterface {
	log.Printf("Proxy address %v -> %v\n", c.Proxy.ProxyAddress, c.Proxy.SiteAddress)

	a := &App{
		conf: c,
	}
	return a
}

type App struct {
	conf AppConfig
}

func (a *App) Run() {
	log.Printf("Application run...")

	go metrics.RunMetrics(a.conf.Profile)

	p := UseCase.NewTmProcessor(a.conf.RunesInWorld)

	myHandler := UseCase.NewHttpProxyTextProcessorHandler(a.conf.Proxy.SiteAddress, p)

	err := UseCase.NewHttpServer().Run(a.conf.Proxy.ProxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
