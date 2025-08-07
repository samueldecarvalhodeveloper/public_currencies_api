package constants

import (
	"os"
	"strings"
)

const SERVER_PORT = ":3000"

func LIST_OF_SERVICES_TO_BE_BALANCED() []string {
	return strings.Split(os.Getenv("LIST_OF_SERVICES_TO_BE_BALANCED"), ",")
}
