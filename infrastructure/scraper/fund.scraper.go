package scraper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/google/uuid"
	corid "github.com/hthl85/aws-lambda-corid"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/fund"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/overview"
)

// FundScraper struct
type FundScraper struct {
	FundJob         *colly.Collector
	HoldingJob      *colly.Collector
	OverviewJob     *colly.Collector
	fundService     *fund.Service
	holdingService  *holding.Service
	overviewService *overview.Service
	log             logger.ContextLog
}

// NewFundScraper create new fund scraper
func NewFundScraper(fs *fund.Service, hs *holding.Service, os *overview.Service, l logger.ContextLog) *FundScraper {
	fj := newScraperJob()
	hj := newScraperJob()
	oj := newScraperJob()

	return &FundScraper{
		FundJob:         fj,
		HoldingJob:      hj,
		OverviewJob:     oj,
		fundService:     fs,
		holdingService:  hs,
		overviewService: os,
		log:             l,
	}
}

// newScraperJob creates a new colly collector with some custom configs
func newScraperJob() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(config.AllowDomain),
		colly.Async(true),
	)

	// Overrides the default timeout (10 seconds) for this collector
	c.SetRequestTimeout(30 * time.Second)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  config.DomainGlob,
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}

// configJobs configs on error handler and on response handler for scaper jobs
func (s *FundScraper) configJobs() {
	s.FundJob.OnError(s.errorHandler)
	s.FundJob.OnResponse(s.processFundListResponse)

	s.HoldingJob.OnError(s.errorHandler)
	s.HoldingJob.OnResponse(s.processFundHoldingResponse)

	s.OverviewJob.OnError(s.errorHandler)
	s.OverviewJob.OnResponse(s.processFundOverviewResponse)
}

// StartJobs start jobs
func (s *FundScraper) StartJobs() {
	ctx := context.Background()

	s.configJobs()

	if err := s.FundJob.Visit(config.FundListURL); err != nil {
		s.log.Error(ctx, "scrape fund list failed", "error", err)
	}

	s.FundJob.Wait()
	s.HoldingJob.Wait()
	s.OverviewJob.Wait()
}

///////////////////////////////////////////////////////////
// Fund List Scraper
///////////////////////////////////////////////////////////

// errorHandler generic error handler for all scaper jobs
func (s *FundScraper) errorHandler(r *colly.Response, err error) {
	ctx := context.Background()
	s.log.Error(ctx, "failed to request url", "url", r.Request.URL, "error", err)
}

func (s *FundScraper) processFundListResponse(r *colly.Response) {
	// create correlation if for processing fund list
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	// define anonymous struct to map fund data from fund list reponse
	d := struct {
		Funds map[string]entities.VanguardFund `json:"fundData,omitempty"`
	}{}

	// unmarshal response data to above struct
	if err := json.Unmarshal(r.Body, &d); err != nil {
		s.log.Error(ctx, "unmarshal fund list response failed", "error", err)
		return
	}

	for key, fund := range d.Funds {
		if fund.Ticker != "" {
			if err := s.fundService.CreateFund(ctx, &fund); err != nil {
				s.log.Error(ctx, "create fund failed", "portId", key, "error", err)
				continue
			}

			// scrape overview data
			overviewURL := config.GetFundOverviewURL(key)
			overviewCTX := colly.NewContext()
			s.OverviewJob.Request("GET", overviewURL, nil, overviewCTX, nil)

			// scrape holding data
			holdingURL := config.GetFundHoldingURL(key, "F", fund.AssetCode)
			holdingCTX := colly.NewContext()
			holdingCTX.Put("portId", key)
			holdingCTX.Put("ticker", fund.Ticker)
			holdingCTX.Put("assetCode", fund.AssetCode)
			s.HoldingJob.Request("GET", holdingURL, nil, holdingCTX, nil)
		}
	}
}

///////////////////////////////////////////////////////////
// Fund Overview Scraper
///////////////////////////////////////////////////////////

func (s *FundScraper) processFundOverviewResponse(r *colly.Response) {
	// create correlation if for processing fund overview
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	overview := &entities.VanguardFundOverview{}
	if err := json.Unmarshal(r.Body, overview); err != nil {
		s.log.Error(ctx, "failed to parse fund overview response", "error", err)
		return
	}

	if err := s.overviewService.CreateOverview(ctx, overview); err != nil {
		s.log.Error(ctx, "failed to create overview", "portId", overview.PortID, "error", err)
	}
}

///////////////////////////////////////////////////////////
// Fund Holding Scraper
///////////////////////////////////////////////////////////

func (s *FundScraper) processFundHoldingResponse(r *colly.Response) {
	// create correlation if for processing fund holding
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	portID := r.Request.Ctx.Get("portId")
	ticker := r.Request.Ctx.Get("ticker")
	assetCode := r.Request.Ctx.Get("assetCode")

	holding := &entities.VanguardFundHolding{
		PortID:    portID,
		Ticker:    ticker,
		AssetCode: assetCode,
	}

	if assetCode == "BOND" {
		if err := json.Unmarshal(r.Body, &holding.Bonds); err != nil {
			s.log.Error(ctx, "failed to parse bond holding response", "error", err)
		}
	} else if assetCode == "EQUITY" {
		if err := json.Unmarshal(r.Body, &holding.Equities); err != nil {
			s.log.Error(ctx, "failed to parse equity holding response", "error", err)
		}
	} else if assetCode == "BALANCED" {
		if err := json.Unmarshal(r.Body, &holding.Balances); err != nil {
			s.log.Error(ctx, "failed to parse balanced holding response", "error", err)
		}
	} else {
		s.log.Error(ctx, "unsupport asset type", "assetCode", assetCode)
		return
	}

	if err := s.holdingService.CreateHolding(ctx, holding); err != nil {
		s.log.Error(ctx, "failed to create holding", "portId", portID, "error", err)
	}
}
