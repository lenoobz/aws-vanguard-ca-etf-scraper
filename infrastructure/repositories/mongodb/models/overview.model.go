package models

import (
	"context"
	"strconv"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OverviewModel represents Vanguard's fund overview details model
type OverviewModel struct {
	ID               *primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	IsActive         bool                    `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt        int64                   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt       int64                   `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Schema           string                  `json:"schema,omitempty" bson:"schema,omitempty"`
	PortID           string                  `json:"portId,omitempty" bson:"portId,omitempty"`
	AssetClass       string                  `json:"assetClass,omitempty" bson:"assetClass,omitempty"`
	Strategy         string                  `json:"strategy,omitempty" bson:"strategy,omitempty"`
	DividendSchedule string                  `json:"dividendSchedule,omitempty" bson:"dividendSchedule,omitempty"`
	ShortName        string                  `json:"shortName,omitempty" bson:"shortName,omitempty"`
	Currency         string                  `json:"baseCurrency,omitempty" bson:"baseCurrency,omitempty"`
	Isin             string                  `json:"isin,omitempty" bson:"isin,omitempty"`
	Sedol            string                  `json:"sedol,omitempty" bson:"sedol,omitempty"`
	Ticker           string                  `json:"ticker,omitempty" bson:"ticker,omitempty"`
	TotalAssets      float64                 `json:"totalAssets,omitempty" bson:"totalAssets,omitempty"`
	Yield12Month     float64                 `json:"yield12Month,omitempty" bson:"yield12Month,omitempty"`
	Price            float64                 `json:"price,omitempty" bson:"price,omitempty"`
	ManagementFee    float64                 `json:"managementFee,omitempty" bson:"managementFee,omitempty"`
	MerValue         float64                 `json:"merValue,omitempty" bson:"merValue,omitempty"`
	SectorWeighting  []*SectorWeightingModel `json:"sectorWeighting,omitempty" bson:"sectorWeighting,omitempty"`
	CountryExposure  []*CountryExposureModel `json:"countryExposure,omitempty" bson:"countryExposure,omitempty"`
}

// SectorWeightingModel is the representation of sector the fund invested
type SectorWeightingModel struct {
	FundPercent float64 `json:"fundPercent,omitempty" bson:"fundPercent,omitempty"`
	LongName    string  `json:"longName,omitempty" bson:"longName,omitempty"`
	SectorType  string  `json:"sectorType,omitempty" bson:"sectorType,omitempty"`
}

// CountryExposureModel is the representation of country the fund exposed
type CountryExposureModel struct {
	CountryName     string  `json:"countryName,omitempty" bson:"countryName,omitempty"`
	HoldingStatCode string  `json:"holdingStatCode,omitempty" bson:"holdingStatCode,omitempty"`
	FundMktPercent  float64 `json:"fundMktPercent,omitempty" bson:"fundMktPercent,omitempty"`
	FundTnaPercent  float64 `json:"fundTnaPercent,omitempty" bson:"fundTnaPercent,omitempty"`
}

// NewOverviewModel create a fund overview model
func NewOverviewModel(ctx context.Context, log logger.IAppLogger, o *entities.Overview) (*OverviewModel, error) {
	var fundOverview = &OverviewModel{}

	if o.PortID != "" {
		fundOverview.PortID = o.PortID
	}

	if o.AssetClass != "" {
		fundOverview.AssetClass = o.AssetClass
	}

	if o.Strategy != "" {
		fundOverview.Strategy = o.Strategy
	}

	if o.DividendSchedule != "" {
		fundOverview.DividendSchedule = o.DividendSchedule
	}

	if o.ShortName != "" {
		fundOverview.ShortName = o.ShortName
	}

	if o.BaseCurrency != "" {
		fundOverview.Currency = o.BaseCurrency
	}

	if o.FundCodesData != nil {
		if o.FundCodesData.Isin != "" {
			fundOverview.Isin = o.FundCodesData.Isin
		}

		if o.FundCodesData.Sedol != "" {
			fundOverview.Sedol = o.FundCodesData.Sedol
		}

		if o.FundCodesData.ExchangeTicker != "" {
			fundOverview.Ticker = o.FundCodesData.ExchangeTicker
		}
	}

	if o.TotalAssets != "" {
		totalAssets, err := strconv.ParseFloat(o.TotalAssets, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.TotalAssets failed", "err", err, "TotalAssets", o.TotalAssets)
			totalAssets = 0
		}

		fundOverview.TotalAssets = totalAssets
	}

	if o.Yield12Month != "" {
		yield12Month, err := strconv.ParseFloat(o.Yield12Month, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.Yield12Month failed", "err", err, "Yield12Month", o.Yield12Month)
			yield12Month = 0
		}

		fundOverview.Yield12Month = yield12Month
	}

	fundOverview.Price = o.Price

	if o.ManagementFee != "" {
		managementFee, err := strconv.ParseFloat(o.ManagementFee, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.ManagementFee failed", "err", err, "ManagementFee", o.ManagementFee)
			managementFee = 0
		}

		fundOverview.ManagementFee = managementFee
	}

	if o.MerValue != "" {
		merValue, err := strconv.ParseFloat(o.MerValue, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.MerValue failed", "err", err, "MerValue", o.MerValue)
			merValue = 0
		}

		fundOverview.MerValue = merValue
	}

	var sectors []*SectorWeightingModel
	for _, v := range o.SectorWeighting {
		sector, err := newSectorWeightingModel(ctx, log, v)

		if err != nil {
			return nil, err
		}

		sectors = append(sectors, sector)
	}
	fundOverview.SectorWeighting = sectors

	var exposures []*CountryExposureModel
	for _, v := range o.CountryExposure {
		exposure, err := newCountryExposureModel(ctx, log, v)

		if err != nil {
			return nil, err
		}

		if exposure != nil {
			exposures = append(exposures, exposure)
		}
	}
	fundOverview.CountryExposure = exposures

	return fundOverview, nil
}

func newSectorWeightingModel(ctx context.Context, log logger.IAppLogger, sector *entities.SectorWeighting) (*SectorWeightingModel, error) {
	var sectorWeighting = &SectorWeightingModel{}

	if sector.LongName != "" {
		sectorWeighting.LongName = sector.LongName
	}

	if sector.SectorType != "" {
		sectorWeighting.SectorType = sector.SectorType
	}

	if sector.FundPercent != "" {
		fundPercent, err := strconv.ParseFloat(sector.FundPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse SectorWeighting.FundPercent failed", "err", err, "FundPercent", sector.FundPercent)
			fundPercent = 0
		}

		sectorWeighting.FundPercent = fundPercent
	}

	return sectorWeighting, nil
}

func newCountryExposureModel(ctx context.Context, log logger.IAppLogger, exposure *entities.CountryExposure) (*CountryExposureModel, error) {
	var countryExposure = &CountryExposureModel{}

	if exposure.CountryName != "" {
		countryExposure.CountryName = exposure.CountryName
	}

	if exposure.HoldingStatCode != "" {
		countryExposure.HoldingStatCode = exposure.HoldingStatCode
	}

	if exposure.FundMktPercent != "" {
		mktPercent, err := strconv.ParseFloat(exposure.FundMktPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse CountryExposure.FundMktPercent failed", "err", err, "FundMktPercent", exposure.FundMktPercent)
			mktPercent = 0
		}

		if mktPercent == 0 {
			return nil, nil
		}

		countryExposure.FundMktPercent = mktPercent
	}

	if exposure.FundTnaPercent != "" {
		tnaPercent, err := strconv.ParseFloat(exposure.FundTnaPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse CountryExposure.FundTnaPercent failed", "err", err, "FundTnaPercent", exposure.FundTnaPercent)
			tnaPercent = 0
		}

		countryExposure.FundTnaPercent = tnaPercent
	}

	return countryExposure, nil
}
