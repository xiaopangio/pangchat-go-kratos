package loadbalance

import (
	"github.com/go-kratos/kratos/v2/registry"
	"math/rand"
)

type LoadBalance interface {
	Pick(instances []*registry.ServiceInstance) *registry.ServiceInstance
}
type RandomLoadBalance struct {
}

func NewRandomLoadBalance() LoadBalance {
	return &RandomLoadBalance{}
}

func (l *RandomLoadBalance) Pick(instances []*registry.ServiceInstance) *registry.ServiceInstance {
	n := len(instances)
	if n == 0 {
		return nil
	}
	return instances[rand.Intn(n)]
}
