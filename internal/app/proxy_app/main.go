package proxy_app

import (
	"habra-tm-habr/internal/app"
	HttpUseCase "habra-tm-habr/internal/pkg/network/http"
	"habra-tm-habr/internal/pkg/network/metrics"
	"habra-tm-habr/internal/pkg/text_processor"
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

	go metrics.RunMetrics(a.conf.ProfileAddress)

	p := text_processor.NewTmProcessor(a.conf.RunesInWorld)
	myHandler := HttpUseCase.NewHttpProxyTextProcessorHandler(a.conf.Proxy.SiteAddress, p)
	srv := HttpUseCase.NewHttpServer()

	err := srv.Run(a.conf.Proxy.ProxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
