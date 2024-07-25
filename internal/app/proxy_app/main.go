package proxy_app

import (
	"habra-tm-habr/config"
	"habra-tm-habr/internal/app"
	HttpUseCase "habra-tm-habr/internal/pkg/network/http"
	"habra-tm-habr/internal/pkg/network/metrics"
	"habra-tm-habr/internal/pkg/text_processor"
	"log"
)

func New(c config.Config) app.ApplicationInterface {
	log.Printf("Proxy address %v -> %v\n", c.ProxyAddress, c.SiteAddress)

	a := &App{
		conf: c,
	}
	return a
}

type App struct {
	conf config.Config
}

func (a *App) Run() {
	log.Printf("Application run...")

	if a.conf.MetricsEnabled {
		go metrics.RunMetrics(a.conf.ProfileAddress)
	}

	p := text_processor.NewTmProcessor(a.conf.RunesInWorld)
	myHandler := HttpUseCase.NewHttpProxyTextProcessorHandler(a.conf.SiteAddress, p)
	srv := HttpUseCase.NewHttpServer()

	err := srv.Run(a.conf.ProxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
