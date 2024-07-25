package main

import (
	"habra-tm-habr/config"
	"habra-tm-habr/internal/app/proxy_app"
	"log"
)

func main() {
	c := config.New()
	a := proxy_app.New(c)
	log.Printf("application: %+v", a)
	log.Printf("config: %+v", c)
	a.Run()
}
