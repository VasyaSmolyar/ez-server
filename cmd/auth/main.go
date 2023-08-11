package main

import (
	"flag"

	"ex-server/internal/auth/server"
)

var (
	configPath = flag.String("conf", "./app/configs/auth.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := server.Init(*configPath)
	if err == nil {
		server.Run()
	}
}
