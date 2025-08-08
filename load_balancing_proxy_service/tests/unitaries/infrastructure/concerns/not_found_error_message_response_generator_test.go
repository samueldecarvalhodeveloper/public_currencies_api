package concerns_test

import (
	"load_balancing_proxy_service/constants"
	"load_balancing_proxy_service/infrastructure/concerns"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfFunctionGetNotFoundErrorMessageResponseReturnsMarshalledNotFoundErrorMessageResponse(
	t *testing.T) {
	notFoundErrorMessageResponse := concerns.GetNotFoundErrorMessageResponse()

	assert.Equal(
		t,
		constants.LIST_OF_BYTES_OF_NOT_FOUND_ERROR_MESSAGE_RESPONSE_BODY(),
		notFoundErrorMessageResponse)
}
