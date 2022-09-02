package services

import (
	"github.com/pkg/errors"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/repositories"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
	"go.uber.org/zap"
)

type DetailsService interface {
	Get(ID int64) (*models.Detail, error)
	Add(req *models.Detail) (p *models.Detail, err error)
}

type DefaultDetailsService struct {
	logger     *zap.Logger
	Repository repositories.DetailsRepository
}

func NewDetailService(logger *zap.Logger, Repository repositories.DetailsRepository) DetailsService {
	return &DefaultDetailsService{
		logger:     logger.With(zap.String("type", "DefaultDetailsService")),
		Repository: Repository,
	}
}

func (s *DefaultDetailsService) Get(ID int64) (p *models.Detail, err error) {
	if p, err = s.Repository.Get(ID); err != nil {
		return nil, errors.Wrap(err, "detail service get detail error")
	}

	return
}

func (s *DefaultDetailsService) Add(req *models.Detail) (p *models.Detail, err error) {

	if p, err = s.Repository.Add(req); err != nil {
		return nil, errors.Wrap(err, "detail service Add detail error")
	}

	return p, err
}
