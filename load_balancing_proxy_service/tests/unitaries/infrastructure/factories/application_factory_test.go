package factories_test

import (
	"github.com/stretchr/testify/assert"

	"load_balancing_proxy_service/infrastructure/factories"
	"testing"
)

func TestIfFunctionGetApplicationInstanceReturnsAnApplicationInstance(
	t *testing.T) {
	instance := factories.GetApplicationInstance()

	assert.NotNil(
		t,
		instance)
}
