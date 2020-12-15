package repositories

import "github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"

// IFundRepository interface
type IFundRepository interface {
	InsertFundOverview(*domains.FundOverview) error
}
