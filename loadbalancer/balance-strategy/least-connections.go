package balancestrategy

import (
	"sync"

	"go-load-balancer/priorityqueue"
)

// LeastConnectionsStrategy
//
//	implements the LoadBalancerStrategy interface.
type LeastConnectionsStrategy struct {
	backends      []string
	servers       map[string]int
	priorityQueue priorityqueue.PriorityQueue
	mu            sync.Mutex
}

// NewLeastConnectionsStrategy
//
//	initializes a least connections strategy with the servers and
//	priority queue.
func NewLeastConnectionsStrategy(backends []string) *LeastConnectionsStrategy {
	priorityQueue := priorityqueue.NewPriorityQueue()
	servers := addServersToMap(backends)

	for _, backend := range backends {
		priorityQueue.Insert(backend, 0)
	}

	return &LeastConnectionsStrategy{
		backends:      backends,
		servers:       servers,
		priorityQueue: *priorityQueue,
	}
}

// addServersToMap
//
//	adds the backends list to a hash map with empty connection counts.
func addServersToMap(backends []string) map[string]int {
	servers := make(map[string]int)
	for _, backend := range backends {
		servers[backend] = 0 // Initialize each server with 0 active connections
	}
	return servers
}

// GetNextBackend
//
//	selects the next backend in a least connections strategy.
func (lc *LeastConnectionsStrategy) GetNextBackend() string {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if len(lc.backends) == 0 {
		return "backends are empty"
	}

	backend, err := lc.priorityQueue.Pop()
	if err != nil {
		// Handle error e@e
		return "error popping priority queue"
	}

	// Increment active connections for the selected server
	server := backend.(string)
	lc.servers[server]++

	// Re-insert the server into the priority queue with updated priority (connection count)
	lc.priorityQueue.Insert(server, float64(lc.servers[server]))

	return server
}

// DecrementConnection
//
//	 should be called when a connection is closed.
//		decrements the connection count for a server.
func (lc *LeastConnectionsStrategy) DecrementConnection(server string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.servers[server]--
	// Update priority in the priority queue
	lc.priorityQueue.UpdatePriority(server, float64(lc.servers[server]))
}
