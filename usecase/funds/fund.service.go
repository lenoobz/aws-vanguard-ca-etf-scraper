package funds

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

// CreateFund creates new fund
func (s *Service) CreateFund(ctx context.Context, fund *entities.Fund) error {
	s.log.Info(ctx, "creating new fund")
	return s.repo.InsertFund(ctx, fund)
}
