package overview

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
func NewService(repo Repo, log logger.ContextLog) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// CreateFundOverview creates new overview
func (s *Service) CreateFundOverview(ctx context.Context, fundOverview *entities.VanguardFundOverview) error {
	s.log.Info(ctx, "create new fund overview")
	return s.repo.InsertFundOverview(ctx, fundOverview)
}
