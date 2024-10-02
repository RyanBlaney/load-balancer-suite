package balancestrategy

import "sync"

// RoundRobinStrategy
//
//	implements the LoadBalancerStrategy interface.
type RoundRobinStrategy struct {
	backends  []string
	currentID int
	mu        sync.Mutex
}

// NewRoundRobinStrategy
//
//	initializes a new RoundRobinStrategy with the ID initialized
//	to 0.
func NewRoundRobinStrategy(backends []string) *RoundRobinStrategy {
	return &RoundRobinStrategy{
		backends:  backends,
		currentID: 0,
	}
}

// GetNextBackend
//
//	selects the next backend in round-robin fashion.
func (rr *RoundRobinStrategy) GetNextBackend() string {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	backend := rr.backends[rr.currentID]
	rr.currentID = (rr.currentID + 1) % len(rr.backends)
	return backend
}
