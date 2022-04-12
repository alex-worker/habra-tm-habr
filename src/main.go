package main

import (
	"habra-tm-habr/src/server"
	"log"
)

const proxyAddress1 = "127.0.0.1:8080"
const siteAddress = "http://habrahabr.ru"

func main() {
	log.Println("Hello world!")
	server.Listen(proxyAddress1, siteAddress)
}
