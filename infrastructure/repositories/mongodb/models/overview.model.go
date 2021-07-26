package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/consts"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/entities"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/utils/datetime"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundOverviewModel struct
type FundOverviewModel struct {
	ID               *primitive.ObjectID      `bson:"_id,omitempty"`
	CreatedAt        int64                    `bson:"createdAt,omitempty"`
	ModifiedAt       int64                    `bson:"modifiedAt,omitempty"`
	Enabled          bool                     `bson:"enabled"`
	Deleted          bool                     `bson:"deleted"`
	Schema           string                   `bson:"schema,omitempty"`
	PortID           string                   `bson:"portId,omitempty"`
	AssetClass       string                   `bson:"assetClass,omitempty"`
	Strategy         string                   `bson:"strategy,omitempty"`
	DividendSchedule string                   `bson:"dividendSchedule,omitempty"`
	Name             string                   `bson:"name,omitempty"`
	ShortName        string                   `bson:"shortName,omitempty"`
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
	DistAmount       float64                  `bson:"distAmount,omitempty"`
	AllocationStock  float64                  `bson:"allocationStock,omitempty"`
	AllocationBond   float64                  `bson:"allocationBond,omitempty"`
	AllocationCash   float64                  `bson:"allocationCash,omitempty"`
	Sectors          []*SectorBreakdownModel  `bson:"sectors,omitempty"`
	Countries        []*CountryBreakdownModel `bson:"countries,omitempty"`
	Dividends        []*DividendHistoryModel  `bson:"dividends,omitempty"`
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
func NewOverviewModel(ctx context.Context, log logger.ContextLog, fundOverview *entities.FundOverview, schemaVersion string) (*FundOverviewModel, error) {
	var fundOverviewModel = &FundOverviewModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if fundOverview.PortID != "" {
		fundOverviewModel.PortID = fundOverview.PortID
	}

	if fundOverview.AssetClass != "" {
		fundOverviewModel.AssetClass = strings.ToUpper(fundOverview.AssetClass)
	}

	if fundOverview.Strategy != "" {
		fundOverviewModel.Strategy = fundOverview.Strategy
	}

	if fundOverview.DividendSchedule != "" {
		fundOverviewModel.DividendSchedule = strings.ToUpper(fundOverview.DividendSchedule)
	}

	if fundOverview.ShortName != "" {
		fundOverviewModel.ShortName = fundOverview.ShortName
	}

	if fundOverview.Name != "" {
		fundOverviewModel.Name = fundOverview.Name
	}

	if fundOverview.BaseCurrency != "" {
		fundOverviewModel.Currency = fundOverview.BaseCurrency
	}

	if fundOverview.FundCode != nil {
		if fundOverview.FundCode.Isin != "" {
			fundOverviewModel.Isin = fundOverview.FundCode.Isin
		}

		if fundOverview.FundCode.Sedol != "" {
			fundOverviewModel.Sedol = fundOverview.FundCode.Sedol
		}

		if fundOverview.FundCode.ExchangeTicker != "" {
			fundOverviewModel.Ticker = ticker.GenYahooTickerFromVanguardTicker(fundOverview.FundCode.ExchangeTicker)
		}
	}

	if fundOverview.TotalAssets != "" {
		totalAssets, err := strconv.ParseFloat(fundOverview.TotalAssets, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.TotalAssets failed", "error", err, "TotalAssets", fundOverview.TotalAssets)
			totalAssets = 0
		}

		fundOverviewModel.TotalAssets = totalAssets
	}

	if fundOverview.Yield12Month != "" {
		yield12Month, err := strconv.ParseFloat(fundOverview.Yield12Month, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.Yield12Month failed", "error", err, "Yield12Month", fundOverview.Yield12Month)
			yield12Month = 0
		}

		fundOverviewModel.Yield12Month = yield12Month
	}

	fundOverviewModel.Price = fundOverview.Price

	if fundOverview.ManagementFee != "" {
		managementFee, err := strconv.ParseFloat(fundOverview.ManagementFee, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.ManagementFee failed", "error", err, "ManagementFee", fundOverview.ManagementFee)
			managementFee = 0
		}

		fundOverviewModel.ManagementFee = managementFee
	}

	if fundOverview.MerFee != "" {
		merFee, err := strconv.ParseFloat(fundOverview.MerFee, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.MerValue failed", "error", err, "MerValue", fundOverview.MerFee)
			merFee = 0
		}

		fundOverviewModel.MerFee = merFee
	}

	if fundOverview.DistYield != "" {
		distYield, err := strconv.ParseFloat(fundOverview.DistYield, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.DistYield failed", "error", err, "DistYield", fundOverview.DistYield)
			distYield = 0
		}

		fundOverviewModel.DistYield = distYield
	}

	if fundOverview.DistAmount != "" {
		distAmount, err := strconv.ParseFloat(fundOverview.DistAmount, 64)

		if err != nil {
			log.Warn(ctx, "parse Overview.DistAmount failed", "error", err, "DistAmount", fundOverview.DistAmount)
			distAmount = 0
		}

		fundOverviewModel.DistAmount = distAmount
	}

	fundOverviewModel.AllocationStock = fundOverview.AllocationStock
	fundOverviewModel.AllocationBond = fundOverview.AllocationBond
	fundOverviewModel.AllocationCash = fundOverview.AllocationCash

	// map sector breakdown model
	var sectorModels []*SectorBreakdownModel
	for _, sector := range fundOverview.Sectors {
		sectorModel, err := newSectorBreakdownModel(ctx, log, sector)

		if err != nil {
			return nil, err
		}

		if sectorModel != nil {
			sectorModels = append(sectorModels, sectorModel)
		}
	}
	fundOverviewModel.Sectors = sectorModels

	// map country breakdown model
	var countryModels []*CountryBreakdownModel
	for _, country := range fundOverview.Countries {
		countryModel, err := newCountryBreakdownModel(ctx, log, country)

		if err != nil {
			return nil, err
		}

		if countryModel != nil {
			countryModels = append(countryModels, countryModel)
		}
	}
	fundOverviewModel.Countries = countryModels

	// map dividend history model
	var divHistoryModels []*DividendHistoryModel
	for _, dividend := range fundOverview.Dividends {
		divHistoryModel, err := newDividendHistoryModel(ctx, log, dividend)

		if err != nil {
			return nil, err
		}

		if divHistoryModel != nil {
			divHistoryModels = append(divHistoryModels, divHistoryModel)
		}
	}
	fundOverviewModel.Dividends = divHistoryModels

	return fundOverviewModel, nil
}

// newSectorBreakdownModel create sector breakdown model
func newSectorBreakdownModel(ctx context.Context, log logger.ContextLog, sectorBreakdown *entities.SectorBreakdown) (*SectorBreakdownModel, error) {
	var sectorBreakdownModel = &SectorBreakdownModel{}

	if sectorBreakdown.SectorName != "" {
		sectorBreakdownModel.SectorName = sectorBreakdown.SectorName

		sectorCode, err := getSectorCode(sectorBreakdown.SectorName)
		if err != nil {
			log.Warn(ctx, "get sector code failed", "error", err, "SectorName", sectorBreakdown.SectorName)
		}

		sectorBreakdownModel.SectorCode = sectorCode
	}

	if sectorBreakdown.FundPercent != "" {
		fundPercent, err := strconv.ParseFloat(sectorBreakdown.FundPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse SectorWeighting.FundPercent failed", "error", err, "FundPercent", sectorBreakdown.FundPercent)
			fundPercent = 0
		}

		sectorBreakdownModel.FundPercent = fundPercent
	}

	// We will try to use BenchmarkPercent if FundPercent is not available
	if sectorBreakdownModel.FundPercent == 0 && sectorBreakdown.BenchmarkPercent != "" {
		benchmarkPercent, err := strconv.ParseFloat(sectorBreakdown.BenchmarkPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse SectorWeighting.BenchmarkPercent failed", "error", err, "BenchmarkPercent", sectorBreakdown.BenchmarkPercent)
			benchmarkPercent = 0
		}

		sectorBreakdownModel.FundPercent = benchmarkPercent
	}

	// We are only interested in sector that has FundPercent/BenchmarkPercent
	if sectorBreakdownModel.FundPercent != 0 {
		return sectorBreakdownModel, nil
	}

	return nil, nil
}

// newCountryBreakdownModel create country breakdown model
func newCountryBreakdownModel(ctx context.Context, log logger.ContextLog, countryBreakdown *entities.CountryBreakdown) (*CountryBreakdownModel, error) {
	var countryBreakdownModel = &CountryBreakdownModel{}

	if countryBreakdown.CountryName != "" {
		countryBreakdownModel.CountryName = countryBreakdown.CountryName

		countryCode, err := getCountryCode(countryBreakdown.CountryName)
		if err != nil {
			log.Warn(ctx, "get country code failed", "error", err, "CountryName", countryBreakdown.CountryName)
		}

		countryBreakdownModel.CountryCode = countryCode
	}

	if countryBreakdown.HoldingStatCode != "" {
		countryBreakdownModel.HoldingStatCode = countryBreakdown.HoldingStatCode
	}

	if countryBreakdown.FundMktPercent != "" {
		fundMktPercent, err := strconv.ParseFloat(countryBreakdown.FundMktPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse CountryExposure.FundMktPercent failed", "error", err, "FundMktPercent", countryBreakdown.FundMktPercent)
			fundMktPercent = 0
		}

		if fundMktPercent == 0 {
			// not interested in fund market with 0 percent
			return nil, nil
		}

		countryBreakdownModel.FundMktPercent = fundMktPercent
	}

	if countryBreakdown.FundTnaPercent != "" {
		fundTnaPercent, err := strconv.ParseFloat(countryBreakdown.FundTnaPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse CountryExposure.FundTnaPercent failed", "error", err, "FundTnaPercent", countryBreakdown.FundTnaPercent)
			fundTnaPercent = 0
		}

		countryBreakdownModel.FundTnaPercent = fundTnaPercent
	}

	return countryBreakdownModel, nil
}

// newDividendHistoryModel create dividend history model
func newDividendHistoryModel(ctx context.Context, log logger.ContextLog, dividendHistories *entities.DividendHistory) (*DividendHistoryModel, error) {
	var dividentHistoryModel = &DividendHistoryModel{}

	dividentHistoryModel.CurrencyCode = dividendHistories.CurrencyCode

	if dividendHistories.Amount != "" {
		amount, err := strconv.ParseFloat(dividendHistories.Amount, 64)
		if err != nil {
			log.Warn(ctx, "parse DistHistory.Amount failed", "error", err, "Amount", dividendHistories.Amount)
			amount = 0
		}

		if amount == 0 {
			// not interested in dividend with 0 amount
			return nil, nil
		}

		dividentHistoryModel.Amount = amount
	}

	if dividendHistories.CurrencyCode != "" {
		asOfDate, err := datetime.GetStarDateFromString(dividendHistories.AsOfDate)
		if err != nil {
			log.Warn(ctx, "parse DistHistory.AsOfDate failed", "error", err, "AsOfDate", dividendHistories.AsOfDate)
			asOfDate = nil
		}

		dividentHistoryModel.AsOfDate = asOfDate
	}

	return dividentHistoryModel, nil
}

// getCountryCode gets country code of a given name
func getCountryCode(name string) (string, error) {
	for _, country := range consts.Countries {
		if strings.EqualFold(strings.ToUpper(country.Name), strings.ToUpper(name)) {
			return country.Alpha3Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find country code for country %s", name)
}

// getSectorCode gets sector code of a given name
func getSectorCode(name string) (string, error) {
	for _, sector := range consts.Sectors {
		if strings.EqualFold(strings.ToUpper(sector.Name), strings.ToUpper(name)) {
			return sector.Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find sector code for sector %s", name)
}
