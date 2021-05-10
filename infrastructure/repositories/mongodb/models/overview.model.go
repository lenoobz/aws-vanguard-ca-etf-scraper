package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/datetime"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardOverviewModel struct
type VanguardOverviewModel struct {
	ID               *primitive.ObjectID      `bson:"_id,omitempty"`
	IsActive         bool                     `bson:"isActive,omitempty"`
	CreatedAt        int64                    `bson:"createdAt,omitempty"`
	ModifiedAt       int64                    `bson:"modifiedAt,omitempty"`
	Schema           string                   `bson:"schema,omitempty"`
	PortID           string                   `bson:"portId,omitempty"`
	AssetClass       string                   `bson:"assetClass,omitempty"`
	Strategy         string                   `bson:"strategy,omitempty"`
	DividendSchedule string                   `bson:"dividendSchedule,omitempty"`
	Name             string                   `bson:"name,omitempty"`
	Currency         string                   `bson:"currency,omitempty"`
	Isin             string                   `bson:"isin,omitempty"`
	Sedol            string                   `bson:"sedol,omitempty"`
	Ticker           string                   `bson:"ticker,omitempty"`
	TotalAssets      float64                  `bson:"totalAssets,omitempty"`
	Yield12Month     float64                  `bson:"yield12Month,omitempty"`
	Price            float64                  `bson:"price,omitempty"`
	ManagementFee    float64                  `bson:"managementFee,omitempty"`
	MerFee           float64                  `bson:"merFee,omitempty"`
	DistYield        float64                  `bson:"distYield,omitempty"`
	AllocationStock  float64                  `bson:"allocationStock,omitempty"`
	AllocationBond   float64                  `bson:"allocationBond,omitempty"`
	AllocationCash   float64                  `bson:"allocationCash,omitempty"`
	Sectors          []*SectorBreakdownModel  `bson:"sectors,omitempty"`
	Countries        []*CountryBreakdownModel `bson:"countries,omitempty"`
	DividendHistory  []*DividendHistoryModel  `bson:"dividendHistory,omitempty"`
}

// SectorBreakdownModel struct
type SectorBreakdownModel struct {
	SectorCode  string  `bson:"sectorCode,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
	FundPercent float64 `bson:"fundPercent,omitempty"`
}

// CountryBreakdownModel struct
type CountryBreakdownModel struct {
	CountryCode     string  `bson:"countryCode,omitempty"`
	CountryName     string  `bson:"countryName,omitempty"`
	FundMktPercent  float64 `bson:"fundMktPercent,omitempty"`
	FundTnaPercent  float64 `bson:"fundTnaPercent,omitempty"`
	HoldingStatCode string  `bson:"holdingStatCode,omitempty"`
}

// DividendHistoryModel struct
type DividendHistoryModel struct {
	Amount       float64    `bson:"amount,omitempty"`
	CurrencyCode string     `bson:"currencyCode,omitempty"`
	AsOfDate     *time.Time `bson:"asOfDate,omitempty"`
}

// NewOverviewModel create a fund overview model
func NewOverviewModel(ctx context.Context, l logger.ContextLog, e *entities.VanguardFundOverview) (*VanguardOverviewModel, error) {
	var m = &VanguardOverviewModel{}

	if e.PortID != "" {
		m.PortID = e.PortID
	}

	if e.AssetClass != "" {
		m.AssetClass = e.AssetClass
	}

	if e.Strategy != "" {
		m.Strategy = e.Strategy
	}

	if e.DividendSchedule != "" {
		m.DividendSchedule = e.DividendSchedule
	}

	if e.ShortName != "" {
		m.Name = e.ShortName
	}

	if e.BaseCurrency != "" {
		m.Currency = e.BaseCurrency
	}

	if e.FundCode != nil {
		if e.FundCode.Isin != "" {
			m.Isin = e.FundCode.Isin
		}

		if e.FundCode.Sedol != "" {
			m.Sedol = e.FundCode.Sedol
		}

		if e.FundCode.ExchangeTicker != "" {
			m.Ticker = ticker.GetYahooTicker(e.FundCode.ExchangeTicker)
		}
	}

	if e.TotalAssets != "" {
		ta, err := strconv.ParseFloat(e.TotalAssets, 64)

		if err != nil {
			l.Warn(ctx, "parse Overview.TotalAssets failed", "error", err, "TotalAssets", e.TotalAssets)
			ta = 0
		}

		m.TotalAssets = ta
	}

	if e.Yield12Month != "" {
		ym, err := strconv.ParseFloat(e.Yield12Month, 64)

		if err != nil {
			l.Warn(ctx, "parse Overview.Yield12Month failed", "error", err, "Yield12Month", e.Yield12Month)
			ym = 0
		}

		m.Yield12Month = ym
	}

	m.Price = e.Price

	if e.ManagementFee != "" {
		v, err := strconv.ParseFloat(e.ManagementFee, 64)

		if err != nil {
			l.Warn(ctx, "parse Overview.ManagementFee failed", "error", err, "ManagementFee", e.ManagementFee)
			v = 0
		}

		m.ManagementFee = v
	}

	if e.MerFee != "" {
		v, err := strconv.ParseFloat(e.MerFee, 64)

		if err != nil {
			l.Warn(ctx, "parse Overview.MerValue failed", "error", err, "MerValue", e.MerFee)
			v = 0
		}

		m.MerFee = v
	}

	if e.DistYield != "" {
		v, err := strconv.ParseFloat(e.DistYield, 64)

		if err != nil {
			l.Warn(ctx, "parse Overview.DistYield failed", "error", err, "DistYield", e.DistYield)
			v = 0
		}

		m.DistYield = v
	}

	m.AllocationStock = e.AllocationStock
	m.AllocationBond = e.AllocationBond
	m.AllocationCash = e.AllocationCash

	// map sector breakdown model
	var sectors []*SectorBreakdownModel
	for _, s := range e.Sectors {
		sector, err := newSectorBreakdownModel(ctx, l, s)

		if err != nil {
			return nil, err
		}

		sectors = append(sectors, sector)
	}
	m.Sectors = sectors

	// map country breakdown model
	var countries []*CountryBreakdownModel
	for _, c := range e.Countries {
		country, err := newCountryBreakdownModel(ctx, l, c)

		if err != nil {
			return nil, err
		}

		if country != nil {
			countries = append(countries, country)
		}
	}
	m.Countries = countries

	// map dividend history model
	var dists []*DividendHistoryModel
	for _, d := range e.DistHistory {
		dist, err := newDividendHistoryModel(ctx, l, d)

		if err != nil {
			return nil, err
		}

		if dist != nil {
			dists = append(dists, dist)
		}
	}
	m.DividendHistory = dists

	return m, nil
}

// newSectorBreakdownModel create sector breakdown model
func newSectorBreakdownModel(ctx context.Context, l logger.ContextLog, e *entities.SectorBreakdown) (*SectorBreakdownModel, error) {
	var m = &SectorBreakdownModel{}

	if e.SectorName != "" {
		m.SectorName = e.SectorName

		sectorCode, err := getSectorCode(e.SectorName)
		if err != nil {
			l.Warn(ctx, "get sector code failed", "error", err, "SectorName", e.SectorName)
		}

		m.SectorCode = sectorCode
	}

	if e.FundPercent != "" {
		v, err := strconv.ParseFloat(e.FundPercent, 64)

		if err != nil {
			l.Warn(ctx, "parse SectorWeighting.FundPercent failed", "error", err, "FundPercent", e.FundPercent)
			v = 0
		}

		m.FundPercent = v
	}

	return m, nil
}

// newCountryBreakdownModel create country breakdown model
func newCountryBreakdownModel(ctx context.Context, l logger.ContextLog, e *entities.CountryBreakdown) (*CountryBreakdownModel, error) {
	var m = &CountryBreakdownModel{}

	if e.CountryName != "" {
		m.CountryName = e.CountryName

		countryCode, err := getCountryCode(e.CountryName)
		if err != nil {
			l.Warn(ctx, "get country code failed", "error", err, "CountryName", e.CountryName)
		}

		m.CountryCode = countryCode
	}

	if e.HoldingStatCode != "" {
		m.HoldingStatCode = e.HoldingStatCode
	}

	if e.FundMktPercent != "" {
		v, err := strconv.ParseFloat(e.FundMktPercent, 64)

		if err != nil {
			l.Warn(ctx, "parse CountryExposure.FundMktPercent failed", "error", err, "FundMktPercent", e.FundMktPercent)
			v = 0
		}

		if v == 0 {
			// not interested in fund market with 0 percent
			return nil, nil
		}

		m.FundMktPercent = v
	}

	if e.FundTnaPercent != "" {
		v, err := strconv.ParseFloat(e.FundTnaPercent, 64)

		if err != nil {
			l.Warn(ctx, "parse CountryExposure.FundTnaPercent failed", "error", err, "FundTnaPercent", e.FundTnaPercent)
			v = 0
		}

		m.FundTnaPercent = v
	}

	return m, nil
}

// newDividendHistoryModel create dividend history model
func newDividendHistoryModel(ctx context.Context, l logger.ContextLog, e *entities.DividendHistory) (*DividendHistoryModel, error) {
	var m = &DividendHistoryModel{}

	m.CurrencyCode = e.CurrencyCode

	if e.Amount != "" {
		v, err := strconv.ParseFloat(e.Amount, 64)
		if err != nil {
			l.Warn(ctx, "parse DistHistory.Amount failed", "error", err, "Amount", e.Amount)
			v = 0
		}

		if v == 0 {
			// not interested in dividend with 0 amount
			return nil, nil
		}

		m.Amount = v
	}

	if e.CurrencyCode != "" {
		v, err := datetime.GetDateStartFromString(e.AsOfDate)
		if err != nil {
			l.Warn(ctx, "parse DistHistory.AsOfDate failed", "error", err, "AsOfDate", e.AsOfDate)
			v = nil
		}

		m.AsOfDate = v
	}

	return m, nil
}

// getCountryCode gets country code of a given name
func getCountryCode(name string) (string, error) {
	for _, v := range consts.Countries {
		if strings.EqualFold(strings.ToUpper(v.Name), strings.ToUpper(name)) {
			return v.Alpha3Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find country code for country %s", name)
}

// getSectorCode gets sector code of a given name
func getSectorCode(name string) (string, error) {
	for _, v := range consts.Sectors {
		if strings.EqualFold(strings.ToUpper(v.Name), strings.ToUpper(name)) {
			return v.Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find sector code for sector %s", name)
}
