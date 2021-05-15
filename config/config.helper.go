package config

import "fmt"

// AllowDomain const
const AllowDomain = "api.vanguard.com"

// DomainGlob const
const DomainGlob = "*vanguard.*"

// FundListURL const
const FundListURL = "https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-listview-data-en.json"

// GetFundOverviewURL get fund overview url
func GetFundOverviewURL(portID string) string {
	return fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-overview-data-etf.json?vars=portId:%s,lang:en&path=[portId=%s][0]", portID, portID)
}

// GetFundDistributionURL get fund distribution url
func GetFundDistributionURL(portID, issueType string) string {
	return fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-price-distribution.json?vars=portId:%s,issueType:%s", portID, issueType)
}

// GetFundHoldingURL get fund holding url
func GetFundHoldingURL(portID, issueType, assetCode string) string {
	var URL string

	switch assetCode {
	case "BOND":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-bond.json?vars=portId:%s,issueType:%s", portID, issueType)
	case "EQUITY":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-equity.json?vars=portId:%s,issueType:%s", portID, issueType)
	case "BALANCED":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-balanced.json?vars=portId:%s,issueType:%s", portID, issueType)
	default:
		break
	}

	return URL
}
