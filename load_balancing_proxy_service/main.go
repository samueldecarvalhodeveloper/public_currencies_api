package main

import (
	"load_balancing_proxy_service/infrastructure/concerns"
)

func main() {

	app := concerns.GetApplicationInstance()

	app.RunServer()
}
