package service

import (
	"botfb/internal/model/entity"
	"context"
)

type IReview interface {
	Create(ctx context.Context, in entity.Review) error
}

var localReview IReview

func Review() IReview {
	if localReview == nil {
		panic("implement not found for interface IReview, forgot register?")
	}
	return localReview
}

func RegisterReview(i IReview) {
	localReview = i
}
