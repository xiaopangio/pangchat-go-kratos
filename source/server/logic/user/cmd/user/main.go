package main

import (
	"flag"
	"os"
	"strings"
	"user/internal/components/registry"

	"user/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
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
	flag.StringVar(&flagconf, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, cf *conf.Bootstrap, gs *grpc.Server, registry *registry.UserRegistry) *kratos.App {
	port := strings.Split(cf.Server.Grpc.Addr, ":")[1]
	return kratos.New(
		kratos.ID(id+"."+port),
		kratos.Name(cf.Server.Name),
		kratos.Version(cf.Server.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(registry),
	)
}

func main() {
	flag.Parse()

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
	logger := log.With(log.NewStdLogger(os.Stdout),
		"time", log.Timestamp("2006-01-02 15:04:05.000"),
		"caller", log.DefaultCaller,
		//"service.id", id,
		//"service.name", bc.Server.Name,
		//"service.version", bc.Server.Version,
		//"trace.id", tracing.TraceID(),
		//"span.id", tracing.SpanID(),
	)
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
