package scrapers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/services"
)

// FundScraper struct
type FundScraper struct {
	FundListCol     *colly.Collector
	FundOverviewCol *colly.Collector
	FundHoldingCol  *colly.Collector
	fundService     services.IFundService
}

// NewFundScraper create new fund scraper
func NewFundScraper(svc services.IFundService) *FundScraper {
	fmt.Println("Create new Fund Scraper")

	fl := newFundListScraper()
	fo := newFundOverviewScraper()
	fh := newFundHoldingScraper()

	fl.OnResponse(processFundListResponse(fo, fh, svc))
	fo.OnResponse(processFundOverviewResponse(svc))
	fh.OnResponse(processFundHoldingResponse(svc))

	return &FundScraper{
		FundListCol:     fl,
		FundOverviewCol: fo,
		FundHoldingCol:  fh,
		fundService:     svc,
	}
}

func newFundListScraper() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(consts.AllowDomain),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return c
}

func newFundOverviewScraper() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(consts.AllowDomain),
		colly.Async(true),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  consts.DomainGlob,
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return c
}

func newFundHoldingScraper() *colly.Collector {
	c := colly.NewCollector(
		// colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains(consts.AllowDomain),
		colly.Async(true),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  consts.DomainGlob,
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return c
}

func processFundListResponse(fo, fh *colly.Collector, svc services.IFundService) colly.ResponseCallback {
	return func(r *colly.Response) {
		rs := &domains.VanguardFunds{}
		if err := json.Unmarshal(r.Body, rs); err != nil {
			fmt.Println("[warning] - error occurred when parsing fund list response", err)
			return
		}

		if err := handleFundList(rs, fo, fh, svc); err != nil {
			fmt.Println("error occurred when processing fund list response", err)
			return
		}
	}
}

func handleFundList(funds *domains.VanguardFunds, fo, fh *colly.Collector, svc services.IFundService) error {
	for key, fund := range funds.IndividualFunds {
		if fund.Ticker != "" {
			if err := svc.CreateIndividualFund(&fund); err != nil {
				return err
			}

			overviewURL := consts.GetFundOverviewURL(key)
			overviewCTX := colly.NewContext()
			fo.Request("GET", overviewURL, nil, overviewCTX, nil)

			holdingURL := consts.GetFundHoldingURL(key, "F", fund.AssetCode)
			holdingCTX := colly.NewContext()
			holdingCTX.Put("portId", key)
			holdingCTX.Put("ticker", fund.Ticker)
			holdingCTX.Put("assetCode", fund.AssetCode)
			fh.Request("GET", holdingURL, nil, holdingCTX, nil)
		}
	}

	return nil
}

func processFundOverviewResponse(svc services.IFundService) colly.ResponseCallback {
	return func(r *colly.Response) {
		rs := &domains.FundOverview{}
		if err := json.Unmarshal(r.Body, rs); err != nil {
			fmt.Println("[warning] - error occurred when parsing fund overview response", err)
			return
		}

		if err := handleFundOverview(rs, svc); err != nil {
			fmt.Println("error occurred when processing fund overview response", err)
			return
		}
	}
}

func handleFundOverview(overview *domains.FundOverview, svc services.IFundService) error {
	if err := svc.CreateFundOverview(overview); err != nil {
		return err
	}

	return nil
}

func processFundHoldingResponse(svc services.IFundService) colly.ResponseCallback {
	return func(r *colly.Response) {
		portID := r.Request.Ctx.Get("portId")
		ticker := r.Request.Ctx.Get("ticker")
		assetCode := r.Request.Ctx.Get("assetCode")

		rs := &domains.FundHolding{
			PortID:    portID,
			Ticker:    ticker,
			AssetCode: assetCode,
		}

		if assetCode == "BOND" {
			if err := json.Unmarshal(r.Body, &rs.BondHolding); err != nil {
				fmt.Println("[warning] - error occurred when parsing bond fund holding response", err)
				return
			}

			if err := handleFundHolding(rs, svc); err != nil {
				fmt.Println("error occurred when processing bond fund holding response", err)
				return
			}
		} else if assetCode == "EQUITY" {
			if err := json.Unmarshal(r.Body, &rs.EquityHolding); err != nil {
				fmt.Println("[warning] - error occurred when parsing equity fund holding response", err)
				return
			}

			if err := handleFundHolding(rs, svc); err != nil {
				fmt.Println("error occurred when processing equity fund holding response", err)
				return
			}
		} else if assetCode == "BALANCED" {
			if err := json.Unmarshal(r.Body, &rs.BalancedHolding); err != nil {
				fmt.Println("[warning] - error occurred when parsing balanced fund holding response", err)
				return
			}

			if err := handleFundHolding(rs, svc); err != nil {
				fmt.Println("error occurred when processing balanced fund holding response", err)
				return
			}
		}
	}
}

func handleFundHolding(fundHolding *domains.FundHolding, svc services.IFundService) error {
	if err := svc.CreateFundHolding(fundHolding); err != nil {
		return err
	}

	return nil
}
