package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	logic "logic/api/v1"
	"logic/internal/components/loadbalance"
	"logic/internal/components/oss"
	"logic/internal/components/redis"
	"logic/internal/components/registry"
	"logic/internal/conf"
	"logic/internal/data/orm/model"
	"logic/pkg"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	// ToolOptionsKey is redis key for tool options
	ToolOptionsKey = "tool_options"
	// EmojisKey is redis key for emojis
	EmojisKey = "emojis"
)

// LogicBiz 不太好分类的业务逻辑类
type LogicBiz struct {
	helper            *log.Helper
	redisCli          *redis.Redis
	connectorRegistry *registry.ConnectorRegistry
	repo              LogicRepo
	lb                loadbalance.LoadBalance
	cf                *conf.Service
	ossClient         *oss.OSSClient
}

// LogicRepo 业务逻辑仓库接口
type LogicRepo interface {
	GetToolOptions(ctx context.Context) (options []*model.ToolOption, err error)
	GetEmojis(ctx context.Context) (emojis []*model.Emoji, err error)
}

// NewLogicBiz 生成业务逻辑类实例
func NewLogicBiz(helper *log.Helper, ossClient *oss.OSSClient, redisCli *redis.Redis, connectorRegistry *registry.ConnectorRegistry, lb loadbalance.LoadBalance, cf *conf.Bootstrap, repo LogicRepo) *LogicBiz {
	return &LogicBiz{helper: helper, ossClient: ossClient, redisCli: redisCli, connectorRegistry: connectorRegistry, lb: lb, cf: cf.Service, repo: repo}
}

// GetConnectorUrl 负载均衡获取connector的url
func (l *LogicBiz) GetConnectorUrl(ctx context.Context) (error error, host, port string) {
	instances, err := l.connectorRegistry.GetService(ctx, l.cf.ConnectorService)
	if err != nil {
		return
	}
	//随机选择一个connector，实现负载均衡
	instance := l.lb.Pick(instances)
	if instance == nil {
		return
	}
	var endpoint string
	for _, s := range instance.Endpoints {
		if strings.Contains(s, "http") {
			endpoint = s
			break
		}
	}
	url := strings.Split(endpoint, "//")[1]
	host = strings.Split(url, ":")[0]
	port = strings.Split(url, ":")[1]
	return nil, host, port
}

// TransferToolOptions 转换工具选项
func (l *LogicBiz) TransferToolOptions(options []*model.ToolOption) []*logic.ToolOption {
	toolOptions := make([]*logic.ToolOption, 0)
	for _, option := range options {
		toolOptions = append(toolOptions, &logic.ToolOption{
			Name: option.Name,
			Icon: option.Icon,
		})
	}
	return toolOptions
}

// TransferToolOption 转换工具选项
func (l *LogicBiz) TransferToolOption(option *model.ToolOption) *logic.ToolOption {
	return &logic.ToolOption{
		Name: option.Name,
		Icon: option.Icon,
	}
}

// GetToolOptions 获取工具选项
func (l *LogicBiz) GetToolOptions(ctx context.Context) ([]*logic.ToolOption, error) {
	if values, err := l.redisCli.SMember(ToolOptionsKey); err != nil {
		l.helper.Error("GetToolOptions from redis ", "err", err)
		return nil, pkg.InternalError("GetToolOptions from redis ", err)
	} else {
		if values == nil || len(values) == 0 {
			//redis中没有数据，从数据库中获取
			res, err := l.repo.GetToolOptions(ctx)
			if err != nil {
				l.helper.Error("GetToolOptions from db ", "err", err)
				return nil, pkg.InternalError("GetToolOptions from db ", err)
			}
			options := l.TransferToolOptions(res)
			//将数据存入redis
			for _, option := range options {
				value, err := json.Marshal(option)
				if err != nil {
					l.helper.Error("GetToolOptions json.Marshal ", "err", err)
					return nil, pkg.InternalError("GetToolOptions json.Marshal ", err)
				}
				if err = l.redisCli.SAdd(ToolOptionsKey, string(value)); err != nil {
					l.helper.Error("GetToolOptions redis.SAdd ", "err", err)
					return nil, pkg.InternalError("GetToolOptions redis.SAdd ", err)
				}
			}
			return options, nil
		} else {
			//	redis中有数据，直接返回
			options := make([]*logic.ToolOption, 0)
			for _, value := range values {
				option := &logic.ToolOption{}
				if err := json.Unmarshal([]byte(value), option); err != nil {
					l.helper.Error("GetToolOptions json.Unmarshal ", "err", err)
					return nil, pkg.InternalError("GetToolOptions json.Unmarshal ", err)
				}
				options = append(options, option)
			}
			return options, nil
		}
	}
}

// TransferEmoji 转换emoji
func (l *LogicBiz) TransferEmoji(emoji *model.Emoji) *logic.Emoji {
	return &logic.Emoji{
		EId:      pkg.FormatInt(emoji.EID),
		EName:    emoji.EName,
		EContent: emoji.EContent,
	}
}

// TransferEmojis 转换emoji
func (l *LogicBiz) TransferEmojis(emojis []*model.Emoji) []*logic.Emoji {
	res := make([]*logic.Emoji, 0)
	for _, emoji := range emojis {
		res = append(res, l.TransferEmoji(emoji))
	}
	return res
}

// GetPreEmojis 获取预置emoji
func (l *LogicBiz) GetPreEmojis(ctx context.Context) ([]*logic.Emoji, error) {
	if values, err := l.redisCli.SMember(EmojisKey); err != nil {
		l.helper.Error("GetPreEmojis from redis ", "err", err)
		return nil, pkg.InternalError("GetPreEmojis from redis ", err)
	} else {
		if values == nil || len(values) == 0 {
			//redis中没有数据，从数据库中获取
			res, err := l.repo.GetEmojis(ctx)
			if err != nil {
				l.helper.Error("GetPreEmojis from db ", "err", err)
				return nil, pkg.InternalError("GetPreEmojis from db ", err)
			}
			emojis := l.TransferEmojis(res)
			//将数据存入redis
			for _, emoji := range emojis {
				value, err := json.Marshal(emoji)
				if err != nil {
					l.helper.Error("GetPreEmojis json.Marshal ", "err", err)
					return nil, pkg.InternalError("GetPreEmojis json.Marshal ", err)
				}
				if err = l.redisCli.SAdd(EmojisKey, string(value)); err != nil {
					l.helper.Error("GetPreEmojis redis.SAdd ", "err", err)
					return nil, pkg.InternalError("GetPreEmojis redis.SAdd ", err)
				}
			}
			return emojis, nil
		} else {
			//	redis中有数据，直接返回
			emojis := make([]*logic.Emoji, 0)
			for _, value := range values {
				emoji := &logic.Emoji{}
				if err := json.Unmarshal([]byte(value), emoji); err != nil {
					l.helper.Error("GetPreEmojis json.Unmarshal ", "err", err)
					return nil, pkg.InternalError("GetPreEmojis json.Unmarshal ", err)
				}
				emojis = append(emojis, emoji)
			}
			return emojis, nil
		}
	}
}

// UploadFile 上传文件
func (l *LogicBiz) UploadFile(stream logic.Logic_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		l.helper.Error("获取文件信息失败", "err:", err.Error())
		return pkg.InternalError("获取文件信息失败")
	}
	fileInfo := req.GetFileInfo()
	fileName := fileInfo.GetName()
	fileSize := fileInfo.GetSize()
	fileType := fileInfo.GetType()
	filePath := fileInfo.GetPath()
	//随机生成文件名
	fileName = fileName + strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63(), 10)
	filePath = filePath + "/" + fileName + "." + fileType
	reader := NewUploadFileReader(stream)
	err = l.ossClient.Bucket.PutObject(filePath, reader)
	if err != nil {
		l.helper.Error("上传文件失败", "err:", err.Error())
		return pkg.InternalError("上传文件失败")
	}
	//将信息返回给客户端
	res := &logic.UploadFileResponse{
		FilePath: filePath,
		FileName: fileName,
		FileSize: fileSize,
		FileType: fileType,
	}
	if err = stream.SendAndClose(res); err != nil {
		l.helper.Error("上传文件失败", "err:", err.Error())
		return pkg.InternalError("上传文件失败")
	}
	return nil
}

func (l *LogicBiz) DownloadFile(filePath string, stream logic.Logic_DownloadFileServer) error {
	reader, err := l.ossClient.Bucket.GetObject(filePath)
	if err != nil {
		l.helper.Error("下载文件失败", "err:", err.Error())
		return pkg.InternalError("下载文件失败")
	}
	defer func(reader io.ReadCloser) {
		err = reader.Close()
		if err != nil {
			l.helper.Error("下载文件失败", "err:", err.Error())
			return
		}
	}(reader)
	buf := make([]byte, 1<<22)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			l.helper.Error("下载文件失败", "err:", err.Error())
			return pkg.InternalError("下载文件失败")
		}
		if n == 0 {
			break
		}
		if err = stream.Send(&logic.DownloadFileResponse{
			ChunkData: buf[:n],
		}); err != nil {
			l.helper.Error("下载文件失败", "err:", err.Error())
			return pkg.InternalError("下载文件失败")
		}
	}
	return nil
}
