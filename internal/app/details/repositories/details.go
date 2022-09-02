package repositories

import (
	"github.com/pkg/errors"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type DetailsRepository interface {
	Get(ID int64) (p *models.Detail, err error)
	Add(req *models.Detail) (p *models.Detail, err error)
}

type MysqlDetailsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlDetailsRepository(logger *zap.Logger, db *gorm.DB) DetailsRepository {
	return &MysqlDetailsRepository{
		logger: logger.With(zap.String("type", "DetailsRepository")),
		db:     db,
	}
}

func (s *MysqlDetailsRepository) Get(ID int64) (p *models.Detail, err error) {
	p = new(models.Detail)
	if err = s.db.Model(p).Where("id = ?", ID).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get product error[id=%d]", ID)
	}
	return
}

func (s *MysqlDetailsRepository) Add(req *models.Detail) (p *models.Detail, err error) {

	req.CreatedTime = time.Now()

	if err = s.db.Create(req).Error; err != nil {
		return nil, errors.Wrapf(err, "add product error[id=%d]", req)
	}

	return req, nil

}
