package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errLog = log.NewHelper(log.DefaultLogger)

func HandlerError(c *gin.Context, err error) error {
	if err == nil {
		return nil
	}
	s, ok := status.FromError(err)
	if !ok {
		errLog.Error(err.Error())
		FailMessage(c, "未知错误")
		return err
	}
	switch s.Code() {
	case codes.Internal:
		errLog.Error(s.Message())
		FailMessage(c, "服务器内部错误")
	case codes.NotFound:
		FailMessage(c, s.Message())
	case codes.InvalidArgument, codes.AlreadyExists:
		FailMessage(c, s.Message())
	default:
		errLog.Error(s.Message())
		FailMessage(c, "未知错误")
	}
	return err
}
