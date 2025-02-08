package main

import (
	"go-kafka-order-producer/config"
	"go-kafka-order-producer/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	s := server.Init()
	s.Run()

}
