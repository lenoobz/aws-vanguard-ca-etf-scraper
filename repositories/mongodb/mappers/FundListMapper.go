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

	if fl.Currency != "" {
		fund.Currency = fl.Currency
	}

	if fl.IssueTypeCode != "" {
		fund.IssueTypeCode = fl.IssueTypeCode
	}

	if fl.PortID != "" {
		fund.PortID = fl.PortID
	}

	if fl.ProductType != "" {
		fund.ProductType = fl.ProductType
	}

	if fl.ManagementFee != "" {
		fund.ManagementFee = fl.ManagementFee
	}

	if fl.MerValue != "" {
		fund.MerValue = fl.MerValue
	}

	return fund, nil
}
