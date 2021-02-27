package fund

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
)

// Service sector
type Service struct {
	repo IFundRepo
	log  logger.IAppLogger
}

// NewFundService create new service
func NewFundService(r IFundRepo, l logger.IAppLogger) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// CreateFund creates new fund
func (s *Service) CreateFund(ctx context.Context, fund *entities.Fund) error {
	s.log.Info(ctx, "create new fund")
	return s.repo.InsertFund(ctx, fund)
}
