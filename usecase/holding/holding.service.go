package holding

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/entities"
)

// Service sector
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(repo Repo, log logger.ContextLog) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// CreateFundHolding creates new fund holding
func (s *Service) CreateFundHolding(ctx context.Context, fundHolding *entities.FundHolding) error {
	s.log.Info(ctx, "create new fund holding")
	return s.repo.InsertFundHolding(ctx, fundHolding)
}
