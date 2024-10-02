
# Go Load Balancer

This project is a simple load balancer implemented in Golang. The load balancer supports different load-balancing strategies such as Round Robin and Least Connections. It uses a priority queue to manage backend servers based on active connections, making it efficient and scalable for distributing traffic among multiple backends.

## Features

- **Multiple Load Balancing Strategies**:
  - **Round Robin**: Distributes traffic evenly by cycling through the list of backend servers.
  - **Least Connections**: Routes traffic to the backend with the fewest active connections, ensuring an even load distribution based on current usage.

- **Thread Safety**: The load balancer uses mutex locks to ensure that the `servers` map and priority queue are safely accessed across multiple goroutines.

- **Extensible**: The load balancer is designed to allow easy implementation of additional load-balancing strategies by implementing the `LoadBalancerStrategy` interface.

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/go-load-balancer.git
   cd go-load-balancer
   ```

2. **Install Dependencies**:

   This project uses Go modules. Simply run the following command to download the dependencies:

   ```bash
   go mod tidy
   ```


### Project Structure

```main.go```: The main entry point for the load balancer. It initializes the backend servers and starts routing traffic using the selected load balancing strategy.

```loadbalancer/```: Contains the core ```LoadBalance``` struct, which is responsible for distributing requests across the backend servers.

```loadbalancer/balance-strategy/```: Contains implementations of different load balancing strategies. The current strategies include:

- Round Robin
- Least Connections

```priorityqueue/```: Contains the implementation of the priority queue, which is used in the Least Connections strategy to keep track of backend servers based on their connection counts.

### Usage
#### 1. Implementing the Load Balancer
To use the load balancer, you need to define the backend servers and the strategy you want to use. In main.go, you can switch between different strategies (e.g., Least Connections, Round Robin).

**Example (Least Connections)**:

    ```go

    package main

    import (
        "fmt"
        "go-load-balancer/loadbalancer"
        "go-load-balancer/loadbalancer/balance-strategy"
    )

    func main() {
        // Define backend servers
        backends := []string{
            "localhost:8081",
            "localhost:8082",
            "localhost:8083",
            "localhost:8084",
            "localhost:8085",
        }

        // Use Least Connections Strategy
        leastConnectionsStrategy := balancestrategy.NewLeastConnectionsStrategy(backends)

        // Initialize the Load Balancer with the chosen strategy
        lb := loadbalancer.NewLoadBalancer(backends, leastConnectionsStrategy)

        // Simulate traffic distribution
        for i := 0; i < len(backends)*2; i++ {
            backend := lb.GetNextBackend()
            fmt.Printf("Redirecting to backend: %s\n", backend)
        }
    }

    ```


#### 2. Adding a New Load Balancing Strategy

To add a new load balancing strategy, implement the LoadBalancerStrategy interface. For example, if you want to add a Random Load Balancer:

1. Define a new struct in the balance-strategy package.
2. Implement the GetNextBackend() method.

```go
package balancestrategy

import (
    "math/rand"
)

// RandomStrategy implements the LoadBalancerStrategy interface
type RandomStrategy struct {
    backends []string
}

// NewRandomStrategy creates a new RandomStrategy
func NewRandomStrategy(backends []string) *RandomStrategy {
    return &RandomStrategy{backends: backends}
}

// GetNextBackend randomly selects a backend
func (r *RandomStrategy) GetNextBackend() string {
    return r.backends[rand.Intn(len(r.backends))]
}
```

Then, you can use this new strategy in your main program.


