package scraper

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/google/uuid"
	corid "github.com/hthl85/aws-lambda-corid"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/distributions"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/funds"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/overview"
)

// FundScraper struct
type FundScraper struct {
	ScrapeFundListJob         *colly.Collector
	ScrapeFundHoldingJob      *colly.Collector
	ScrapeFundOverviewJob     *colly.Collector
	ScrapeFundDistributionJob *colly.Collector
	fundService               *funds.Service
	holdingService            *holding.Service
	overviewService           *overview.Service
	distributionService       *distributions.Service
	log                       logger.ContextLog
}

// NewFundScraper create new fund scraper
func NewFundScraper(fundService *funds.Service, holdingService *holding.Service, overviewService *overview.Service, distributionService *distributions.Service, log logger.ContextLog) *FundScraper {
	scrapeFundListJob := newScraperJob()
	scrapeFundHoldingJob := newScraperJob()
	scrapeFundOverviewJob := newScraperJob()
	scrapeFundDistributionJob := newScraperJob()

	return &FundScraper{
		ScrapeFundListJob:         scrapeFundListJob,
		ScrapeFundHoldingJob:      scrapeFundHoldingJob,
		ScrapeFundOverviewJob:     scrapeFundOverviewJob,
		ScrapeFundDistributionJob: scrapeFundDistributionJob,
		fundService:               fundService,
		holdingService:            holdingService,
		overviewService:           overviewService,
		distributionService:       distributionService,
		log:                       log,
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
		RandomDelay: 2 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}

// configJobs configs on error handler and on response handler for scaper jobs
func (s *FundScraper) configJobs() {
	s.ScrapeFundListJob.OnError(s.errorHandler)
	s.ScrapeFundListJob.OnResponse(s.processFundListResponse)

	s.ScrapeFundHoldingJob.OnError(s.errorHandler)
	s.ScrapeFundHoldingJob.OnResponse(s.processFundHoldingResponse)

	s.ScrapeFundOverviewJob.OnError(s.errorHandler)
	s.ScrapeFundOverviewJob.OnResponse(s.processFundOverviewResponse)

	s.ScrapeFundDistributionJob.OnError(s.errorHandler)
	s.ScrapeFundDistributionJob.OnResponse(s.processFundDistributionResponse)
}

// ScrapeAllVanguardFundsDetails scrape all Vanguard funds details
func (s *FundScraper) ScrapeAllVanguardFundsDetails() {
	ctx := context.Background()

	s.configJobs()

	if err := s.ScrapeFundListJob.Visit(config.FundListURL); err != nil {
		s.log.Error(ctx, "scrape fund list failed", "error", err)
	}

	s.ScrapeFundListJob.Wait()
	s.ScrapeFundHoldingJob.Wait()
	s.ScrapeFundOverviewJob.Wait()
	s.ScrapeFundDistributionJob.Wait()
}

// errorHandler generic error handler for all scaper jobs
func (s *FundScraper) errorHandler(r *colly.Response, err error) {
	ctx := context.Background()
	s.log.Error(ctx, "failed to request url", "url", r.Request.URL, "error", err)
}

///////////////////////////////////////////////////////////
// Fund List Scraper
///////////////////////////////////////////////////////////

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
			s.ScrapeFundOverviewJob.Request("GET", overviewURL, nil, overviewCTX, nil)

			// scrape holding data
			holdingURL := config.GetFundHoldingURL(key, "F", fund.AssetCode)
			holdingCTX := colly.NewContext()
			holdingCTX.Put("portId", key)
			holdingCTX.Put("ticker", fund.Ticker)
			holdingCTX.Put("assetCode", fund.AssetCode)
			s.ScrapeFundHoldingJob.Request("GET", holdingURL, nil, holdingCTX, nil)

			// scrape distribution data
			distributionURL := config.GetFundDistributionURL(key, "F")
			distributionCTX := colly.NewContext()
			distributionCTX.Put("portId", key)
			distributionCTX.Put("ticker", fund.Ticker)
			s.ScrapeFundDistributionJob.Request("GET", distributionURL, nil, distributionCTX, nil)
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

	if err := s.overviewService.CreateFundOverview(ctx, overview); err != nil {
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

	if strings.EqualFold(assetCode, consts.BOND) {
		if err := json.Unmarshal(r.Body, &holding.Bonds); err != nil {
			s.log.Error(ctx, "failed to parse bond holding response", "error", err)
		}
	} else if strings.EqualFold(assetCode, consts.EQUITY) {
		if err := json.Unmarshal(r.Body, &holding.Equities); err != nil {
			s.log.Error(ctx, "failed to parse equity holding response", "error", err)
		}
	} else if strings.EqualFold(assetCode, consts.BALANCED) {
		if err := json.Unmarshal(r.Body, &holding.Balances); err != nil {
			s.log.Error(ctx, "failed to parse balanced holding response", "error", err)
		}
	} else {
		s.log.Error(ctx, "unsupport asset type", "assetCode", assetCode)
		return
	}

	if err := s.holdingService.CreateFundHolding(ctx, holding); err != nil {
		s.log.Error(ctx, "failed to create holding", "portId", portID, "ticker", ticker, "error", err)
	}
}

///////////////////////////////////////////////////////////
// Fund Distribution Scraper
///////////////////////////////////////////////////////////

func (s *FundScraper) processFundDistributionResponse(r *colly.Response) {
	// create correlation if for processing fund holding
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	portID := r.Request.Ctx.Get("portId")
	ticker := r.Request.Ctx.Get("ticker")

	fundDistribution := &entities.VanguardFundDistribution{}

	if err := json.Unmarshal(r.Body, &fundDistribution); err != nil {
		s.log.Error(ctx, "failed to parse fund distribution response", "error", err)
	}

	fundDistribution.DistributionDetails.Ticker = ticker
	if err := s.distributionService.CreateFundDistribution(ctx, fundDistribution); err != nil {
		s.log.Error(ctx, "failed to create fund distribution", "portId", portID, "ticker", ticker, "error", err)
	}
}
