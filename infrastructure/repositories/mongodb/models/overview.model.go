package models

import (
	"context"
	"strconv"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardOverviewModel represents Vanguard fund overview model
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
	Sectors          []*SectorBreakdownModel  `bson:"sectors,omitempty"`
	Countries        []*CountryBreakdownModel `bson:"countries,omitempty"`
}

// SectorBreakdownModel is the representation of sector the fund invested
type SectorBreakdownModel struct {
	FundPercent float64 `bson:"fundPercent,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
}

// CountryBreakdownModel is the representation of country the fund exposed
type CountryBreakdownModel struct {
	CountryName     string  `bson:"countryName,omitempty"`
	FundMktPercent  float64 `bson:"fundMktPercent,omitempty"`
	FundTnaPercent  float64 `bson:"fundTnaPercent,omitempty"`
	HoldingStatCode string  `bson:"holdingStatCode,omitempty"`
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
			m.Ticker = e.FundCode.ExchangeTicker
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

	var sectors []*SectorBreakdownModel
	for _, s := range e.Sectors {
		sector, err := newSectorBreakdownModel(ctx, l, s)

		if err != nil {
			return nil, err
		}

		sectors = append(sectors, sector)
	}
	m.Sectors = sectors

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

	return m, nil
}

func newSectorBreakdownModel(ctx context.Context, l logger.ContextLog, e *entities.SectorBreakdown) (*SectorBreakdownModel, error) {
	var m = &SectorBreakdownModel{}

	if e.SectorName != "" {
		m.SectorName = e.SectorName
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

func newCountryBreakdownModel(ctx context.Context, l logger.ContextLog, e *entities.CountryBreakdown) (*CountryBreakdownModel, error) {
	var m = &CountryBreakdownModel{}

	if e.CountryName != "" {
		m.CountryName = e.CountryName
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
