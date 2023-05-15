package logger

import "github.com/go-kratos/kratos/v2/log"

func NewHelper(logger log.Logger) *log.Helper {
	return log.NewHelper(logger)
}
