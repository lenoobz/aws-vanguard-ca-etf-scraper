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

// CreateIndividualFund creates fund
func (svc *FundService) CreateIndividualFund(f *domains.IndividualFund) error {
	return svc.fundRepo.InsertIndividualFund(f)
}

// CreateFundOverview creates fund overview
func (svc *FundService) CreateFundOverview(fo *domains.FundOverview) error {
	return svc.fundRepo.InsertFundOverview(fo)
}

// CreateFundHolding creates fund overview
func (svc *FundService) CreateFundHolding(fh *domains.FundHolding) error {
	return svc.fundRepo.InsertFundHolding(fh)
}
