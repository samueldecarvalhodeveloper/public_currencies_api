package loadbalancer

import specifications "load_balancing_proxy_service/infrastructure/specifications/domains"

type LoadBalancer struct {
	currentServerIndex int

	listOfServersToBeBalanced []string
}

func New(listOfServersToBeBalanced []string) LoadBalancer {
	return LoadBalancer{
		listOfServersToBeBalanced: listOfServersToBeBalanced,
	}
}

func (loadBalancer *LoadBalancer) GetCurrentServer() string {
	currentServer := loadBalancer.listOfServersToBeBalanced[loadBalancer.currentServerIndex]

	if specifications.IsCurrentServerTheLast(loadBalancer.currentServerIndex, loadBalancer.listOfServersToBeBalanced) {
		loadBalancer.currentServerIndex = 0
	} else {
		loadBalancer.currentServerIndex++
	}

	return currentServer
}
