package services

import "github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"

// IFundService is the interface that wraps the basic to  method.
type IFundService interface {
	CreateFundOverview(*domains.FundOverview) error
}
