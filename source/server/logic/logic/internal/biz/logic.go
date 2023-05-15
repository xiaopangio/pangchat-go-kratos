package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"logic/internal/components/loadbalance"
	"logic/internal/components/redis"
	"logic/internal/components/registry"
	"logic/internal/conf"
	"strings"
)

type LogicBiz struct {
	helper            *log.Helper
	redisCli          *redis.Redis
	connectorRegistry *registry.ConnectorRegistry
	lb                loadbalance.LoadBalance
	cf                *conf.Service
}

func NewLogicBiz(helper *log.Helper, redisCli *redis.Redis, connectorRegistry *registry.ConnectorRegistry, lb loadbalance.LoadBalance, cf *conf.Service) *LogicBiz {
	return &LogicBiz{helper: helper, redisCli: redisCli, connectorRegistry: connectorRegistry, lb: lb, cf: cf}
}
func (l *LogicBiz) GetConnectorUrl(ctx context.Context) (error error, host, port string) {
	instances, err := l.connectorRegistry.GetService(ctx, l.cf.ConnectorService)
	l.helper.Info("GetConnectorUrl", "instances", instances)
	if err != nil {
		return
	}
	//随机选择一个connector，实现负载均衡
	instance := l.lb.Pick(instances)
	if instance == nil {
		return
	}
	l.helper.Info("GetConnectorUrl", "instance", instance.Endpoints)
	url := strings.Split(instance.Endpoints[1], "//")[1]
	host = strings.Split(url, ":")[0]
	port = strings.Split(url, ":")[1]
	return nil, host, port
}
