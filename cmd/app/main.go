package main

import (
	"habra-tm-habr/config"
	"habra-tm-habr/internal/app"
	"log"
)

func main() {
	c := config.New()
	a := app.New(c)
	log.Printf("application: %+v", a)
	log.Printf("config: %+v", c)
	a.Run()
}
