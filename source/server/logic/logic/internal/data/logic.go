package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"logic/internal/biz"
	"logic/internal/data/orm/dal"
	"logic/internal/data/orm/model"
	"logic/pkg"
)

type LogicRepoImpl struct {
	Data   *Data
	helper *log.Helper
}

func (l *LogicRepoImpl) GetEmojis(ctx context.Context) (emojis []*model.Emoji, err error) {
	if err = pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	emojis, err = dal.Emoji.WithContext(ctx).Find()
	if err != nil {
		l.helper.Errorf("get emojis from db error: %v", err)
		return nil, pkg.InternalError("get emojis from db", err)
	}
	return
}

func NewLogicRepoImpl(data *Data, helper *log.Helper) biz.LogicRepo {
	return &LogicRepoImpl{Data: data, helper: helper}
}

func (l *LogicRepoImpl) GetToolOptions(ctx context.Context) ([]*model.ToolOption, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	options, err := dal.ToolOption.WithContext(ctx).Find()
	if err != nil {
		l.helper.Errorf("get tool options from db error: %v", err)
		return nil, pkg.InternalError("get tool options from db", err)
	}
	return options, nil
}
