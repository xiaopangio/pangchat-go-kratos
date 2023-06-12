package main

import (
	"flag"
	cfg "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"strconv"
	"strings"

	"api-gateway/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

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

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

// NewLogger new a logger.
func NewLogger(bc *conf.Bootstrap) log.Logger {
	//判断是否有日志文件夹，没有则创建
	_, err := os.Stat("../logs")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("../logs", os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
	logFilePath := "../logs/" + bc.Server.Name + "_" + strings.Split(bc.Server.Http.Addr, ":")[1] + ".log"
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.With(log.NewStdLogger(logFile),
		"time", log.Timestamp("2006-01-02 15:04:05.000"),
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", bc.Server.Name,
		"service.version", bc.Server.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	return logger
}

// NewConfig new a config config.
func NewConfig() *conf.Bootstrap {
	//加载本地配置文件
	bc := LocalLoadConfig()
	bc = RemoteLoadConfig(bc)
	return bc
}
func LocalLoadConfig() *conf.Bootstrap {
	//加载本地配置文件
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer func(c config.Config) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return &bc
}
func RemoteLoadConfig(bc *conf.Bootstrap) *conf.Bootstrap {
	//加载远程配置
	addrs := bc.Nacos.Addrs
	if len(addrs) > 0 {
		addr := addrs[0]
		host := strings.Split(addr, ":")[0]
		port, err := strconv.Atoi(strings.Split(addr, ":")[1])
		if err != nil {
			panic(err)
		}
		sc := []constant.ServerConfig{*constant.NewServerConfig(host, uint64(port))}
		cc := constant.ClientConfig{
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogLevel:            "debug",
		}
		client, err := clients.NewConfigClient(vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
		if err != nil {
			panic(err)
		}
		println(bc.Config.Group, bc.Config.DataId)
		source := cfg.NewConfigSource(client, cfg.WithGroup(bc.Config.Group), cfg.WithDataID(bc.Config.DataId))
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf), source,
			),
		)
		defer func(c config.Config) {
			err := c.Close()
			if err != nil {
				panic(err)
			}
		}(c)
		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(bc); err != nil {
			panic(err)
		}
	}
	return bc
}
func main() {
	flag.Parse()
	bc := NewConfig()
	logger := NewLogger(bc)
	app, cleanup, err := wireApp(bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
