package main

import (
	"flag"

	"ex-server/internal/object/server"
)

var (
	configPath = flag.String("conf", "./app/configs/object.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := server.Init(*configPath)
	if err == nil {
		server.Run()
	}
}
