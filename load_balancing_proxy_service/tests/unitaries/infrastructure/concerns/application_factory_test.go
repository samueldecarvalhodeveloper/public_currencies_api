package concerns_test

import (
	"github.com/stretchr/testify/assert"

	"load_balancing_proxy_service/infrastructure/concerns"
	"testing"
)

func TestIfFunctionGetApplicationInstanceReturnsAnApplicationInstance(
	t *testing.T) {
	instance := concerns.GetApplicationInstance()

	assert.NotNil(
		t,
		instance)
}
