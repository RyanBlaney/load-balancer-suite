package main

import (
	"fmt"

	"go-load-balancer/loadbalancer"
	"go-load-balancer/loadbalancer/balance-strategy"
)

func main() {
	backends := []string{
		"localhost:8081",
		"localhost:8082",
		"localhost:8083",
		"localhost:8084",
		"localhost:8085",
	}

	roundRobinStrategy := balancestrategy.NewLeastConnectionsStrategy(backends)

	lb := loadbalancer.NewLoadBalancer(backends, roundRobinStrategy)

	for i := 0; i < len(backends)*2; i++ {
		backend := lb.GetNextBackend()
		fmt.Printf("Redirecting to backend: %s\n", backend)
	}
}
