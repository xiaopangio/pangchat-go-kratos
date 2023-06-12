package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"net/url"
	"online/internal/components/redis"
	"online/internal/components/registry"
	"online/internal/conf"
	"os"
	"strings"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, cf *conf.Bootstrap, registry *registry.OnlineRegistry, endpoints []*url.URL, grpcSrv *grpc.Server, redis *redis.Redis) *kratos.App {
	server := cf.Server
	port := strings.Split(server.Grpc.Addr, ":")[1]
	return kratos.New(
		kratos.ID(id+"."+port),
		kratos.Name(server.Name),
		kratos.Version(server.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(registry),
		kratos.Server(
			grpcSrv,
		),
		kratos.Endpoint(endpoints...),
		kratos.AfterStop(redis.ResetOnlineDevices),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.Timestamp("2006-01-02 15:04:05.000"),
		"caller", log.DefaultCaller,
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
