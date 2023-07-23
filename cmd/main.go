package main

import (
	"ex-server/internal/handler"
	"ex-server/internal/server"
)

const serverPort = 8080

func main() {
	server.Init(serverPort, handler.Init()).Run()
}
