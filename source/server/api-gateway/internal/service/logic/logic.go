package service_logic

import (
	"api-gateway/api/v1/logic/logic"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"mime"
	"mime/multipart"
	"strings"
)

// LogicService is a service for the Logic API.
type LogicService struct {
	client logic.LogicClient
	helper *log.Helper
	Jwt    *auth.JwtManager
}

// NewLogicService creates a new Logic service.
func NewLogicService(client logic.LogicClient, helper *log.Helper, jwt *auth.JwtManager) *LogicService {
	return &LogicService{client: client, helper: helper, Jwt: jwt}
}

// GetConnectorUrl 获取连接器地址
func (l *LogicService) GetConnectorUrl(ctx *gin.Context) {
	reply, err := l.client.GetConnectorUrl(ctx, &logic.GetConnectorUrlRequest{})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, reply)
}

// GetToolOptions 获取工具选项
func (l *LogicService) GetToolOptions(ctx *gin.Context) {
	reply, err := l.client.GetToolOptions(ctx, &logic.GetToolOptionsRequest{})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, reply)
}

// GetPreEmojis 获取预设表情
func (l *LogicService) GetPreEmojis(ctx *gin.Context) {
	reply, err := l.client.GetPreEmojis(ctx, &logic.GetPreEmojisRequest{})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, reply)
}

// UploadFile 上传文件
func (l *LogicService) UploadFile(ctx *gin.Context) {
	c, cancel := pkg.NewContext(ctx)
	defer cancel()
	stream, err := l.client.UploadFile(c)
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		e := fmt.Sprintf("UploadFile:"+"FormFile err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	f, err := file.Open()
	if err != nil {
		e := fmt.Sprintf("UploadFile:"+"Open err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			e := fmt.Sprintf("UploadFile:"+"Close err: %v", err)
			l.helper.Errorf(e)
			pkg.FailMessage(ctx, e)
			return
		}
	}(f)
	splits := strings.Split(file.Filename, ".")
	if len(splits) < 2 {
		e := fmt.Sprintf("UploadFile:"+"file name err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	fileName := splits[0]
	fileType := splits[1]
	fileSize := file.Size
	err = stream.Send(&logic.UploadFileRequest{
		Data: &logic.UploadFileRequest_FileInfo{
			FileInfo: &logic.FileInfo{
				Name: fileName,
				Path: fileType,
				Size: fileSize,
				Type: fileType,
			},
		},
	})
	if err != nil {
		e := fmt.Sprintf("UploadFile:"+"Send fileInfo err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	buf := make([]byte, 1<<22)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			e := fmt.Sprintf("UploadFile:"+"Read err: %v", err)
			l.helper.Errorf(e)
			pkg.FailMessage(ctx, e)
			return
		}
		err = stream.Send(&logic.UploadFileRequest{
			Data: &logic.UploadFileRequest_ChunkData{
				ChunkData: buf[:n],
			},
		})
		if err != nil {
			e := fmt.Sprintf("UploadFile:"+"Send chunkData err: %v", err)
			l.helper.Errorf(e)
			pkg.FailMessage(ctx, e)
			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil && err != io.EOF {
		e := fmt.Sprintf("UploadFile:"+"CloseAndRecv err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	pkg.Ok(ctx, reply)
}

// DownloadFile 下载文件
func (l *LogicService) DownloadFile(ctx *gin.Context) {
	var req DownloadFileRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		e := fmt.Sprintf("DownloadFile:"+"ShouldBindQuery err: %v", err)
		l.helper.Errorf(e)
		pkg.FailMessage(ctx, e)
		return
	}
	c, cancel := pkg.NewContext(ctx)
	defer cancel()
	stream, err := l.client.DownloadFile(c, &logic.DownloadFileRequest{
		FilePath: req.FilePath,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+req.FilePath)
	fileType := strings.Split(req.FilePath, ".")[1]
	//	将filetype转为mime
	ctx.Header("Content-Transfer-Encoding", mime.TypeByExtension("."+fileType))
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			e := fmt.Sprintf("DownloadFile:"+"Recv err: %v", err)
			l.helper.Errorf(e)
			pkg.FailMessage(ctx, e)
			return
		}
		_, err = ctx.Writer.Write(response.GetChunkData())
		if err != nil {
			e := fmt.Sprintf("DownloadFile:"+"Write err: %v", err)
			l.helper.Errorf(e)
			pkg.FailMessage(ctx, e)
			return
		}
	}
}
