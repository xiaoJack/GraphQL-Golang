package services

import (
	"context"
	"github.com/xiaoJack/GraphQL-Golang/api/proto"
	"go.uber.org/zap"
)

type GraphService interface {
	DetailGet(ctx context.Context, ID int64) (*proto.Detail, error)
	DetailAdd(ctx context.Context, name string, Price float32) (p *proto.Detail, err error)
}

type DefaultGraphServiceService struct {
	logger     *zap.Logger
	detailsSvc proto.DetailsClient
	ratingsSvc proto.RatingsClient
	reviewsSvc proto.ReviewsClient
}

func NewDetailService(logger *zap.Logger, detailsSvc proto.DetailsClient, ratingsSvc proto.RatingsClient, reviewsSvc proto.ReviewsClient) GraphService {

	return &DefaultGraphServiceService{
		logger:     logger.With(zap.String("type", "DefaultDetailsService")),
		detailsSvc: detailsSvc,
		ratingsSvc: ratingsSvc,
		reviewsSvc: reviewsSvc,
	}
}

func (s *DefaultGraphServiceService) DetailGet(ctx context.Context, ID int64) (p *proto.Detail, err error) {

	return s.detailsSvc.Get(ctx, &proto.GetDetailRequest{Id: ID})
}

func (s *DefaultGraphServiceService) DetailAdd(ctx context.Context, name string, Price float32) (p *proto.Detail, err error) {

	return s.detailsSvc.Add(ctx, &proto.Detail{Name: name, Price: Price})
}
