package fund

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

// Service sector
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(r Repo, l logger.ContextLog) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// CreateFund creates new fund
func (s *Service) CreateFund(ctx context.Context, e *entities.VanguardFund) error {
	s.log.Info(ctx, "creating new fund")
	return s.repo.InsertFund(ctx, e)
}
