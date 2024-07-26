package main

import (
	"habra-tm-habr/internal/app/proxy_app"
	"log"
)

func main() {
	c := proxy_app.GetDefaultConfig()
	c.RunesInWorld = 6

	a := proxy_app.New(c)
	log.Printf("application: %+v", a)
	log.Printf("config: %+v", c)
	a.Run()
}
