package holding

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

// CreateHolding creates new holding
func (s *Service) CreateHolding(ctx context.Context, e *entities.VanguardFundHolding) error {
	s.log.Info(ctx, "create new holding")
	return s.repo.InsertHolding(ctx, e)
}
