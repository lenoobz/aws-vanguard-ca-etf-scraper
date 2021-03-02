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
func NewService(r Repo, l logger.ContextLog) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// CreateOverview creates new overview
func (s *Service) CreateOverview(ctx context.Context, e *entities.VanguardFundOverview) error {
	s.log.Info(ctx, "create new overview")
	return s.repo.InsertOverview(ctx, e)
}
