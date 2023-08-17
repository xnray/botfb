package review

import (
	"botfb/internal/dao"
	"botfb/internal/model/entity"
	"botfb/internal/service"
	"context"
)

func init() {
	service.RegisterReview(New())
}

type review struct {
}

func (r review) Create(ctx context.Context, in entity.Review) error {
	_, err := dao.Review.Ctx(ctx).Data(in).OmitEmpty().Save()
	return err
}

func New() *review {
	return &review{}
}
