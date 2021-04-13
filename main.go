package main

import (
	"os"
	"os/signal"
	"syscall"
	"bro-code/golang-tutorials/server"
)

func main() {
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
}