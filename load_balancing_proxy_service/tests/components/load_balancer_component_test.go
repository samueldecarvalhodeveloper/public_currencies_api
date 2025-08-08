package components_test

import (
	"load_balancing_proxy_service/constants"
	loadbalancer "load_balancing_proxy_service/domains/load_balancer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfLoadBalancerReturnsCurrentServersToBeUsedOnRequestsLoadBalancing(
	t *testing.T) {
	loadbalancerInstance :=
		loadbalancer.New(
			constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())

	for currentServerIndex := 0; currentServerIndex < len(constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK())+1; currentServerIndex++ {
		currentServer := loadbalancerInstance.GetCurrentServer()

		if currentServerIndex == len(constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK()) {
			assert.Equal(
				t,
				constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK()[0],
				currentServer)
		} else {
			assert.Equal(
				t,
				constants.LIST_OF_SERVICES_TO_BE_BALANCED_MOCK()[currentServerIndex],
				currentServer)
		}
	}

}
