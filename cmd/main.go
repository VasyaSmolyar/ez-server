package main

import (
	"ex-server/internal/server"
	"flag"
)

var (
	configPath = flag.String("conf", "./configs/app.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := server.Init(*configPath)
	if err == nil {
		server.Run()
	}
}
