package main

import (
	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/infrastructure/factories"
)

func main() {
	app := factories.GetApplicationInstance()

	app.RunServer(constants.SERVER_PORT())
}
