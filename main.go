package main

import (
	"go-tcp-server/server"
	"log"
)

func main() {
	tcpListener := server.New(&server.Config{Host: "127.0.0.1", Port: "8124"})
	log.Println("Server running")
	tcpListener.Run()
}
