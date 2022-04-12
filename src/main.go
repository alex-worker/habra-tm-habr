package main

import (
	"habra-tm-habr/src/server"
	"log"
)

const proxyAddress1 = "127.0.0.1:8080"
const siteAddress = "http://habrahabr.ru"

func main() {
	log.Println("Hello world!")

	//done := make(chan os. Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//go server.Listen(proxyAddress1, siteAddress)
	server.Listen(proxyAddress1, siteAddress)

	//<-done
	//log.Println("Have a nice day!")
}
