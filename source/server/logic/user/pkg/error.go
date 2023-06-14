package pkg

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ContextErr(c context.Context) error {
	switch c.Err() {
	case context.Canceled:
		return status.Errorf(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Errorf(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
func IsNotRecordNotFoundError(err error) error {
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
func FailedPreconditionError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.FailedPrecondition, msg)
	}
	return status.Errorf(codes.FailedPrecondition, msg, a...)
}
func InternalError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.Internal, msg)
	}
	return status.Errorf(codes.Internal, msg, a)
}
func AlreadyExistsError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.AlreadyExists, msg)
	}
	return status.Errorf(codes.AlreadyExists, msg, a)
}
func NotFoundError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.NotFound, msg)
	}
	return status.Errorf(codes.NotFound, msg, a)
}
func PermissionDeniedError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.PermissionDenied, msg)
	}
	return status.Errorf(codes.PermissionDenied, msg, a)
}
func InvalidArgumentError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.InvalidArgument, msg)
	}
	return status.Errorf(codes.InvalidArgument, msg, a)
}
func UnauthenticatedError(msg string, a ...any) error {
	if a == nil {
		return status.Errorf(codes.Unauthenticated, msg)
	}
	return status.Errorf(codes.Unauthenticated, msg, a)
}
func JsonUnmarshalError(err error) error {
	return fmt.Errorf("解析json失败: %s", err.Error())
}
