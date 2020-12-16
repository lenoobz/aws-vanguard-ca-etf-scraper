package services

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories"
)

// FundService struct
type FundService struct {
	fundRepo repositories.IFundRepository
}

// NewFundService create as new service
func NewFundService(fundRepo repositories.IFundRepository) *FundService {
	return &FundService{
		fundRepo: fundRepo,
	}
}

// CreateFundOverview creates fund overview
func (svc *FundService) CreateFundOverview(fo *domains.FundOverview) error {
	return svc.fundRepo.InsertFundOverview(fo)
}
