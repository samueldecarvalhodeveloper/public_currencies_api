package main

import "load_balancing_proxy_service/application"

func main() {
	app := application.New()

	app.RunServer()
}
