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

func LIST_OF_SERVICES_TO_BE_BALANCED_MOCK() []string {
	return []string{"first_server", "second_server", "third_server"}
}

const REQUEST_METHOD = "GET"

const NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY = "{\"message\":\"Not Found!\"}"
