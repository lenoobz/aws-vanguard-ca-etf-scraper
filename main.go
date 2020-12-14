package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

type fundList struct {
	FundData map[string]fundData `json:"fundData,omitempty"`
}

type fundData struct {
	Ticker             string `json:"TICKER,omitempty"`
	AssetCode          string `json:"assetCode,omitempty"`
	ManagementTypeCode string `json:"managementTypeCode,omitempty"`
	ManagementFee      string `json:"managementFee,omitempty"`
	MerValue           string `json:"merValue,omitempty"`
	ShortName          string `json:"shortName,omitempty"`
	Currency           string `json:"currency,omitempty"`
}

type portfolioOverview struct {
	PortID           string             `json:"portId,omitempty"`
	TotalAssets      string             `json:"totalAssets,omitempty"`
	DividendSchedule string             `json:"dividendSchedule,omitempty"`
	ShortName        string             `json:"shortName,omitempty"`
	Yield12Month     string             `json:"yield12Month,omitempty"`
	Price            float64            `json:"price,omitempty"`
	SectorWeighting  []*sectorWeighting `json:"sectorWeighting,omitempty"`
	CountryExposure  []*countryExposure `json:"countryExposure,omitempty"`
}

type sectorWeighting struct {
	FundPercent string `json:"fundPercent,omitempty"`
	LongName    string `json:"longName,omitempty"`
	SectorType  string `json:"sectorType,omitempty"`
}

type countryExposure struct {
	CountryName     string `json:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty"`
}

type holdingDetail struct {
	SectorWeightStock []*sectorWeightStock `json:"sectorWeightStock,omitempty"`
	SectorWeightBond  []*sectorWeightBond  `json:"sectorWeightBond,omitempty"`
}

type sectorWeightStock struct {
	Holding          string  `json:"holding,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
	Type             string  `json:"type,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty"`
	Currency         string  `json:"currency,omitempty"`
}

type sectorWeightBond struct {
	Securities       string  `json:"securities,omitempty"`
	Type             string  `json:"type,omitempty"`
	Rate             float64 `json:"rate,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
}

var baseSearchURL = "https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-listview-data-en.json"

func portfolioOverviewURL(portID string) string {
	return fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-overview-data-etf.json?vars=portId:%s,lang:en&path=[portId=%s][0]", portID, portID)
}

func holdingDetailsURL(portID, issueType, assetCode string) string {
	var URL string

	switch assetCode {
	case "BOND":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-bond.jsonp?vars=portId:%s,issueType:%s", portID, issueType)
		break
	case "EQUITY":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-equity.json?vars=portId:%s,issueType:%s", portID, issueType)
		break
	case "BALANCED":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-balanced.json?vars=portId:%s,issueType:%s", portID, issueType)
		break
	default:
		break
	}

	return URL
}

func main() {
	fundListCol := colly.NewCollector(
		// Attach a debugger to the collector
		colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("api.vanguard.com"),
		// colly.Async(true),
	)

	extensions.RandomUserAgent(fundListCol)
	extensions.Referer(fundListCol)

	portfolioOverviewCol := fundListCol.Clone()
	extensions.RandomUserAgent(portfolioOverviewCol)
	extensions.Referer(portfolioOverviewCol)
	portfolioOverviewCol.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	holdingDetailsCol := fundListCol.Clone()
	extensions.RandomUserAgent(holdingDetailsCol)
	extensions.Referer(holdingDetailsCol)
	holdingDetailsCol.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	fundListCol.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	fundListCol.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing response...")
		rs := &fundList{}
		err := json.Unmarshal(r.Body, rs)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// counter := 0
		for key, element := range rs.FundData {
			overviewURL := portfolioOverviewURL(key)
			holdingURL := holdingDetailsURL(key, "F", element.AssetCode)
			portfolioOverviewCol.Request("GET", overviewURL, nil, nil, nil)
			holdingDetailsCol.Request("GET", holdingURL, nil, nil, nil)
		}
	})

	// Set error handler
	fundListCol.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	fundListCol.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Before making a request print "Visiting ..."
	portfolioOverviewCol.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	portfolioOverviewCol.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing response...")
		rs := &portfolioOverview{}
		err := json.Unmarshal(r.Body, rs)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Price is ", rs.Price)
	})

	// Set error handler
	portfolioOverviewCol.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	portfolioOverviewCol.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Before making a request print "Visiting ..."
	holdingDetailsCol.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	holdingDetailsCol.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing response...")
		var rs []*holdingDetail

		// rs := &holdingDetails
		err := json.Unmarshal(r.Body, &rs)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("There are total ", len(rs[0].SectorWeightBond))
		fmt.Println("There are total ", len(rs[0].SectorWeightStock))
	})

	// Set error handler
	holdingDetailsCol.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	holdingDetailsCol.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	if err := fundListCol.Visit(baseSearchURL); err != nil {
		fmt.Println("Error:", err)
	}

	// Wait until threads are finished
	// holdingDetailsCol.Wait()
	// portfolioOverviewCol.Wait()
	// fundListCol.Wait()
}
