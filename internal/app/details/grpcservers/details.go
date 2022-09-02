package grpcservers

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/xiaoJack/GraphQL-Golang/api/proto"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/services"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DetailsServer struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsServer(logger *zap.Logger, ps services.DetailsService) (*DetailsServer, error) {
	return &DetailsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *proto.GetDetailRequest) (*proto.Detail, error) {
	p, err := s.service.Get(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "details grpc service get detail error")
	}
	ct, err := ptypes.TimestampProto(p.CreatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "convert create time error")
	}

	resp := &proto.Detail{
		Id:          p.ID,
		Name:        p.Name,
		Price:       float32(p.Price),
		CreatedTime: ct,
	}

	return resp, nil
}

func (s *DetailsServer) Add(ctx context.Context, req *proto.Detail) (*proto.Detail, error) {
	p, err := s.service.Add(&models.Detail{Name: req.Name, Price: float64(req.Price)})

	s.logger.Debug("Add", zap.Any("p", p), zap.Any("err", err))
	if err != nil {
		return nil, errors.Wrap(err, "details grpc service get detail error")
	}

	resp := &proto.Detail{
		Id:          p.ID,
		Name:        p.Name,
		Price:       float32(p.Price),
		CreatedTime: timestamppb.New(p.CreatedTime),
	}

	return resp, nil
}
