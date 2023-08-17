package service

import (
	"botfb/internal/model/entity"
	"context"
)

type IMessageTemplate interface {
	Create(ctx context.Context, in entity.MessageTemplate) error
	View(ctx context.Context, id string) (entity.MessageTemplate, error)
}

var localMessageTemplate IMessageTemplate

func MessageTemplate() IMessageTemplate {
	if localMessageTemplate == nil {
		panic("implement not found for interface IMessageTemplate , forgot register?")
	}
	return localMessageTemplate
}

func RegisterMessageTemplate(i IMessageTemplate) {
	localMessageTemplate = i
}
