package mappers

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/models"
)

// MapIndividualFundDomain converts fund domain to model
func MapIndividualFundDomain(fl *domains.IndividualFund) (*models.IndividualFundModel, error) {
	var fund = &models.IndividualFundModel{}

	if fl.Ticker != "" {
		fund.Ticker = fl.Ticker
	}

	if fl.Name != "" {
		fund.Name = fl.Name
	}

	if fl.AssetCode != "" {
		fund.AssetCode = fl.AssetCode
	}

	return fund, nil
}
