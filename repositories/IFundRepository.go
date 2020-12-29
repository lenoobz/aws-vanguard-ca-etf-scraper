package repositories

import "github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"

// IFundRepository interface
type IFundRepository interface {
	InsertIndividualFund(*domains.IndividualFund) error
	InsertFundOverview(*domains.FundOverview) error
	InsertFundHolding(*domains.FundHolding) error
}
