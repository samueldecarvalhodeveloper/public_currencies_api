package main

import (
	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/infrastructure/concerns"
)

func main() {
	app := concerns.GetApplicationInstance()

	app.RunServer(constants.SERVER_PORT())
}
