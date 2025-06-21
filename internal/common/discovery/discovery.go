package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {
	Register(ctx context.Context, instanceId, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceId, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceId, serviceName string) error
}

func GenerateInstanceId(serviceName string) string {
	// Generate a unique instance ID based on service name and host port
	x := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	return fmt.Sprintf("%s-%d", serviceName, x)
}
