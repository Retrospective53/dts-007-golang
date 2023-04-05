package main

import (
	"remakech7/server"

	_ "github.com/lib/pq"
)

func main() {
	server.NewHttpServer()
}