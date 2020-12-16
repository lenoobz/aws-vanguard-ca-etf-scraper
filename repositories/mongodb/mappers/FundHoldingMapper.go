package mappers

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/models"
)

// MapFundHoldingDomain converts fund holding domain to model
func MapFundHoldingDomain(fund *domains.FundHolding) (*models.FundHoldingModel, error) {
	var fundHolding = &models.FundHoldingModel{}

	if fund.AssetCode == "BOND" {

		return fundHolding, nil
	}

	return nil, nil
}
