package distributions

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

// CreateFundDistribution creates new fund distribution
func (s *Service) CreateFundDistribution(ctx context.Context, fundDistribution *entities.VanguardFundDistribution) error {
	s.log.Info(ctx, "create new fund distribution")
	return s.repo.InsertFundDistribution(ctx, fundDistribution)
}
