package main

import (
	"ex-server/internal/server"
)

const (
	serverPort = 8080
	configPath = "./configs/app.json"
)

func main() {
	server, err := server.Init(serverPort, configPath)
	if err == nil {
		server.Run()
	}
}
