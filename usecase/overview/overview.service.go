package overview

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
)

// Service sector
type Service struct {
	repo IOverviewRepo
	log  logger.IAppLogger
}

// NewOverviewService create new service
func NewOverviewService(r IOverviewRepo, l logger.IAppLogger) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// CreateOverview creates new overview
func (s *Service) CreateOverview(ctx context.Context, overview *entities.Overview) error {
	s.log.Info(ctx, "create new overview")
	return s.repo.InsertOverview(ctx, overview)
}
