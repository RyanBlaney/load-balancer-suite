package balancestrategy

// LoadBalancerStrategy
//
//		defines the interface for selecting a backend.
//	 Allows for various configurations of load balancing.
type LoadBalancerStrategy interface {
	GetNextBackend() string
}
