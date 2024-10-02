package loadbalancer

import (
	"sync"

	balancestrategy "go-load-balancer/loadbalancer/balance-strategy"
)

// LoadBalancer
//
//	 is the data structure encapsulating the load balancing
//		logic for distributing load between servers.
type LoadBalancer struct {
	backends []string
	mu       sync.Mutex
	strategy balancestrategy.LoadBalancerStrategy
}

// NewLoadBalancer
//
//	initializes a new load balancer containing the slice of
//	backends addresses to manage.
func NewLoadBalancer(
	backends []string,
	strategy balancestrategy.LoadBalancerStrategy,
) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
		strategy: strategy,
	}
}

// GetNextBackend
//
//	calls the strategy's GetNextBackend method to select a backend.
func (lb *LoadBalancer) GetNextBackend() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	return lb.strategy.GetNextBackend()
}
