package specifications_test

import (
	"github.com/stretchr/testify/assert"

	"load_balancing_proxy_service/constants"
	specifications "load_balancing_proxy_service/infrastructure/specifications/domains"
	"testing"
)

func TestIfFunctionIsCurrentServerTheLastReturnsTrueIfCurrentServerIsTheLastOfTheList(
	t *testing.T) {
	currentServerIsTheLast := specifications.IsCurrentServerTheLast(
		len(constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())-1,
		constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())
	currentServerIsNotTheLast := specifications.IsCurrentServerTheLast(
		0,
		constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())

	assert.True(
		t,
		currentServerIsTheLast)
	assert.False(
		t,
		currentServerIsNotTheLast)
}
