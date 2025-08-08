package constants

import (
	"fmt"
	"os"
	"strings"
	"time"

	"strconv"
)

func SERVER_PORT() int {
	serverPort := os.Getenv("SERVER_PORT")

	serverPortNumber, _ := strconv.Atoi(serverPort)

	return serverPortNumber
}

func LIST_OF_SERVICES_TO_BE_BALANCED() []string {
	return strings.Split(os.Getenv("LIST_OF_SERVICES_TO_BE_BALANCED"), ",")
}

func LIST_OF_SERVICES_TO_BE_BALANCED_MOCK() []string {
	return []string{"first_server", "second_server", "third_server"}
}

const REQUEST_METHOD = "GET"

const NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY = "{\"message\":\"Not Found!\"}"

const LOCAL_HOST_PORT = 3000

func LOCAL_HOST_URL() string {
	return fmt.Sprintf("http://localhost:%d", LOCAL_HOST_PORT)
}

const SERVER_START_UP_DELAY = 1 * time.Second
