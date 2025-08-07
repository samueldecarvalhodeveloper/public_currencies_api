package concerns_test

import (
	"load_balancing_proxy_service/infrastructure/concerns"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfMethodGetApplicationInstanceReturnsAnApplicationInstance(
	t *testing.T) {
	instance := concerns.GetApplicationInstance()

	assert.NotNil(
		t,
		instance)
}
