package main

import (
	"flag"

	"ex-server/internal/task/server"
)

var (
	configPath = flag.String("conf", "./configs/task.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := server.Init(*configPath)
	if err == nil {
		server.Run()
	}
}
