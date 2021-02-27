package holding

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
)

// Service sector
type Service struct {
	repo IHoldingRepo
	log  logger.IAppLogger
}

// NewHoldingService create new service
func NewHoldingService(r IHoldingRepo, l logger.IAppLogger) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// CreateHolding creates new holding
func (s *Service) CreateHolding(ctx context.Context, holding *entities.Holding) error {
	s.log.Info(ctx, "create new holding")
	return s.repo.InsertHolding(ctx, holding)
}
