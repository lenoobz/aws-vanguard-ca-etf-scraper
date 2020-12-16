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
	fundService     services.IFundService
}

// NewFundScraper create new fund scraper
func NewFundScraper(svc services.IFundService) *FundScraper {
	fl := newFundListScraper()
	fo := newFundOverviewScraper()

	fl.OnResponse(processFundListResponse(fo))
	fo.OnResponse(processFundOverviewResponse(svc))

	return &FundScraper{
		FundListCol:     fl,
		FundOverviewCol: fo,
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

func processFundListResponse(fo *colly.Collector) colly.ResponseCallback {
	return func(r *colly.Response) {
		rs := &domains.VanguardFunds{}
		if err := json.Unmarshal(r.Body, rs); err != nil {
			fmt.Println("error occurred when parsing fund list response", err)
			return
		}

		if err := handleFundList(rs, fo); err != nil {
			fmt.Println("error occurred when processing fund list response", err)
			return
		}
	}
}

func handleFundList(funds *domains.VanguardFunds, fo *colly.Collector) error {
	for key, element := range funds.IndividualFunds {
		if element.Ticker != "" {
			// fmt.Println("Key:", key, "=>", "Ticker:", element.Ticker, "=>", "Asset Code:", element.AssetCode)

			overviewURL := consts.GetFundOverviewURL(key)
			if err := fo.Visit(overviewURL); err != nil {
				fmt.Println("error occurred when visit url", overviewURL, err)
			}
		}
	}

	return nil
}

func processFundOverviewResponse(svc services.IFundService) colly.ResponseCallback {
	return func(r *colly.Response) {
		rs := &domains.FundOverview{}
		if err := json.Unmarshal(r.Body, rs); err != nil {
			fmt.Println("error occurred when parsing fund overview response", err)
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
