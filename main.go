package main

import "challenge-6/routers"

func main() {
	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}