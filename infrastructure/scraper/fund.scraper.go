package scraper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/google/uuid"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/fund"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/overview"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/corid"
)

// FundScraper struct
type FundScraper struct {
	FundJob         *colly.Collector
	HoldingJob      *colly.Collector
	OverviewJob     *colly.Collector
	log             logger.IAppLogger
	fundService     fund.IFundService
	holdingService  holding.IHoldingService
	overviewService overview.IOverviewService
}

// NewFundScraper create new fund scraper
func NewFundScraper(fundSvc fund.IFundService, holdingSvc holding.IHoldingService, overviewSvc overview.IOverviewService, l logger.IAppLogger) *FundScraper {
	fundJob := newScraperJob()
	holdingJob := newScraperJob()
	overviewJob := newScraperJob()

	return &FundScraper{
		FundJob:         fundJob,
		HoldingJob:      holdingJob,
		OverviewJob:     overviewJob,
		fundService:     fundSvc,
		holdingService:  holdingSvc,
		overviewService: overviewSvc,
		log:             l,
	}
}

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
		s.log.Error(ctx, "scrape fund list failed", "err", err)
	}

	s.FundJob.Wait()
	s.HoldingJob.Wait()
	s.OverviewJob.Wait()
}

///////////////////////////////////////////////////////////
// Fund List Scraper
///////////////////////////////////////////////////////////

func (s *FundScraper) errorHandler(r *colly.Response, err error) {
	ctx := context.Background()
	s.log.Error(ctx, "failed to request url", "url", r.Request.URL, "error", err)
}

func (s *FundScraper) processFundListResponse(r *colly.Response) {
	// create correlation if for processing fund list
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	funds := &entities.VanguardFunds{}
	if err := json.Unmarshal(r.Body, funds); err != nil {
		s.log.Warn(ctx, "failed to parse fund list response", "err", err)
		return
	}

	for key, fund := range funds.IndividualFunds {
		if fund.Ticker != "" {
			if err := s.fundService.CreateFund(ctx, &fund); err != nil {
				s.log.Warn(ctx, "failed to create fund", "portId", key)
				continue
			}

			overviewURL := config.GetFundOverviewURL(key)
			overviewCTX := colly.NewContext()
			s.OverviewJob.Request("GET", overviewURL, nil, overviewCTX, nil)

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

	overview := &entities.Overview{}
	if err := json.Unmarshal(r.Body, overview); err != nil {
		s.log.Warn(ctx, "failed to parse fund overview response", "err", err)
		return
	}

	if err := s.overviewService.CreateOverview(ctx, overview); err != nil {
		s.log.Warn(ctx, "failed to create overview", "portId", overview.PortID)
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

	holding := &entities.Holding{
		PortID:    portID,
		Ticker:    ticker,
		AssetCode: assetCode,
	}

	if assetCode == "BOND" {
		if err := json.Unmarshal(r.Body, &holding.BondHolding); err != nil {
			s.log.Warn(ctx, "failed to parse bond holding response", "err", err)
		}
	} else if assetCode == "EQUITY" {
		if err := json.Unmarshal(r.Body, &holding.EquityHolding); err != nil {
			s.log.Warn(ctx, "failed to parse equity holding response", "err", err)
		}
	} else if assetCode == "BALANCED" {
		if err := json.Unmarshal(r.Body, &holding.BalancedHolding); err != nil {
			s.log.Warn(ctx, "failed to parse balanced holding response", "err", err)
		}
	} else {
		s.log.Error(ctx, "unsupport asset type", "assetCode", assetCode)
		return
	}

	if err := s.holdingService.CreateHolding(ctx, holding); err != nil {
		s.log.Warn(ctx, "failed to create holding", "portId", portID)
	}
}
