package main

import (
	"ex-server/internal/server"
)

const (
	configPath = "./configs/app.json"
)

func main() {
	server, err := server.Init(configPath)
	if err == nil {
		server.Run()
	}
}
