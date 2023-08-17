package message_template

import (
	"botfb/internal/dao"
	"botfb/internal/model/entity"
	"botfb/internal/service"
	"context"
)

func init() {
	service.RegisterMessageTemplate(New())
}

type messageTemplate struct {
}

func (r *messageTemplate) View(ctx context.Context, id string) (entity.MessageTemplate, error) {
	var out entity.MessageTemplate
	if err := dao.MessageTemplate.Ctx(ctx).Where(dao.MessageTemplate.Columns().TemplateId, id).Scan(&out); err != nil {
		return out, err
	}
	return out, nil
}

func (r *messageTemplate) Create(ctx context.Context, in entity.MessageTemplate) error {
	_, err := dao.MessageTemplate.Ctx(ctx).Data(in).OmitEmpty().Save()
	return err
}

func New() *messageTemplate {
	return &messageTemplate{}
}
