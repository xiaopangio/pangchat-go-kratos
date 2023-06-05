package data

import (
	"context"
	"logic/internal/biz"
	"logic/internal/data/orm/dal"
	"logic/internal/data/orm/model"
	"logic/pkg"
)

type LogicRepoImpl struct {
	Data *Data
}

func (l *LogicRepoImpl) GetEmojis(ctx context.Context) (emojis []*model.Emoji, err error) {
	if err = pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	emojis, err = dal.Emoji.WithContext(ctx).Find()
	if err != nil {
		return nil, pkg.InternalError("get emojis from db", err)
	}
	return
}

func NewLogicRepoImpl(data *Data) biz.LogicRepo {
	return &LogicRepoImpl{Data: data}
}

func (l *LogicRepoImpl) GetToolOptions(ctx context.Context) ([]*model.ToolOption, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	options, err := dal.ToolOption.WithContext(ctx).Find()
	if err != nil {
		return nil, pkg.InternalError("get tool options from db", err)
	}
	return options, nil
}
