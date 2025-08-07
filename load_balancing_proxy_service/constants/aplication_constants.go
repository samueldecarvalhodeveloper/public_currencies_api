package constants

import (
	"fmt"
	"os"
	"strings"
)

func SERVER_PORT() string {
	serverPort := os.Getenv("SERVER_PORT")

	return fmt.Sprintf(":%s", serverPort)
}

func LIST_OF_SERVICES_TO_BE_BALANCED() []string {
	return strings.Split(os.Getenv("LIST_OF_SERVICES_TO_BE_BALANCED"), ",")
}
