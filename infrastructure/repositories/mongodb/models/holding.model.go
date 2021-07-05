package models

import (
	"context"
	"strings"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundHoldingModel struct
type FundHoldingModel struct {
	ID         *primitive.ObjectID       `bson:"_id,omitempty"`
	CreatedAt  int64                     `bson:"createdAt,omitempty"`
	ModifiedAt int64                     `bson:"modifiedAt,omitempty"`
	Enabled    bool                      `bson:"enabled"`
	Deleted    bool                      `bson:"deleted"`
	Schema     string                    `bson:"schema,omitempty"`
	PortID     string                    `bson:"portId,omitempty"`
	Ticker     string                    `bson:"ticker,omitempty"`
	AssetCode  string                    `bson:"assetCode,omitempty"`
	Bonds      []*SectorWeightBondModel  `bson:"bondHolding,omitempty"`
	Stocks     []*SectorWeightStockModel `bson:"stockHolding,omitempty"`
}

// SectorWeightBondModel struct
type SectorWeightBondModel struct {
	FaceAmount       float64 `bson:"faceAmount,omitempty"`
	MarketValPercent float64 `bson:"marketValPercent,omitempty"`
	MarketValue      float64 `bson:"marketValue,omitempty"`
	Rate             float64 `bson:"rate,omitempty"`
	Type             string  `bson:"type,omitempty"`
}

// SectorWeightStockModel struct
type SectorWeightStockModel struct {
	MarketValPercent float64 `bson:"marketValPercent,omitempty"`
	MarketValue      float64 `bson:"marketValue,omitempty"`
	Shares           float64 `bson:"shares,omitempty"`
	Symbol           string  `bson:"symbol,omitempty"`
	Type             string  `bson:"type,omitempty"`
}

// NewFundHoldingModel create a fund holding model
func NewFundHoldingModel(ctx context.Context, log logger.ContextLog, fundHolding *entities.FundHolding, schemaVersion string) (*FundHoldingModel, error) {
	if strings.EqualFold(fundHolding.AssetCode, consts.BOND) {
		return newBondHolding(ctx, log, fundHolding, schemaVersion)
	}

	if strings.EqualFold(fundHolding.AssetCode, consts.EQUITY) {
		return newEquityHolding(ctx, log, fundHolding, schemaVersion)
	}

	if strings.EqualFold(fundHolding.AssetCode, consts.BALANCED) {
		return newBalanceHolding(ctx, log, fundHolding, schemaVersion)
	}

	return nil, nil
}

func newBondHolding(ctx context.Context, log logger.ContextLog, fundHolding *entities.FundHolding, schemaVersion string) (*FundHoldingModel, error) {
	var fundHoldingModel = &FundHoldingModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if fundHolding.PortID != "" {
		fundHoldingModel.PortID = fundHolding.PortID
	}

	if fundHolding.Ticker != "" {
		fundHoldingModel.Ticker = ticker.GenYahooTickerFromVanguardTicker(fundHolding.Ticker)
	}

	fundHoldingModel.AssetCode = fundHolding.AssetCode

	if len(fundHolding.Bonds) == 1 {
		var sectorWeightBondModels []*SectorWeightBondModel
		holding := fundHolding.Bonds[0]

		for _, sectorWeightBond := range holding.SectorWeightBonds {
			sectorWeightBondModel, err := newSectorWeightBondModel(ctx, log, sectorWeightBond)
			if err != nil {
				continue
			}

			sectorWeightBondModels = append(sectorWeightBondModels, sectorWeightBondModel)
		}

		fundHoldingModel.Bonds = sectorWeightBondModels
	}

	return fundHoldingModel, nil
}

func newEquityHolding(ctx context.Context, log logger.ContextLog, fundHolding *entities.FundHolding, schemaVersion string) (*FundHoldingModel, error) {
	var fundHoldingModel = &FundHoldingModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if fundHolding.PortID != "" {
		fundHoldingModel.PortID = fundHolding.PortID
	}

	if fundHolding.Ticker != "" {
		fundHoldingModel.Ticker = ticker.GenYahooTickerFromVanguardTicker(fundHolding.Ticker)
	}

	fundHoldingModel.AssetCode = fundHolding.AssetCode

	if len(fundHolding.Equities) == 1 {
		var sectorWeightStockModels []*SectorWeightStockModel
		holding := fundHolding.Equities[0]

		for _, sectorWeightStock := range holding.SectorWeightStocks {
			sectorWeightStockModel, err := newSectorWeightStockModel(ctx, log, sectorWeightStock)
			if err != nil {
				continue
			}

			sectorWeightStockModels = append(sectorWeightStockModels, sectorWeightStockModel)
		}

		fundHoldingModel.Stocks = sectorWeightStockModels
	}

	return fundHoldingModel, nil
}

func newBalanceHolding(ctx context.Context, log logger.ContextLog, fundHolding *entities.FundHolding, schemaVersion string) (*FundHoldingModel, error) {
	var fundHoldingModel = &FundHoldingModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if fundHolding.PortID != "" {
		fundHoldingModel.PortID = fundHolding.PortID
	}

	if fundHolding.Ticker != "" {
		fundHoldingModel.Ticker = ticker.GenYahooTickerFromVanguardTicker(fundHolding.Ticker)
	}

	fundHoldingModel.AssetCode = fundHolding.AssetCode

	if len(fundHolding.Balances) == 1 {
		var sectorWeightStockModels []*SectorWeightStockModel
		var sectorWeightBondModels []*SectorWeightBondModel
		holding := fundHolding.Balances[0]

		for _, sectorWeightStock := range holding.SectorWeightStocks {
			sectorWeightStockModel, err := newSectorWeightStockModel(ctx, log, sectorWeightStock)
			if err != nil {
				continue
			}

			sectorWeightStockModels = append(sectorWeightStockModels, sectorWeightStockModel)
		}

		for _, sectorWeightBond := range holding.SectorWeightBonds {
			sectorWeightBondModel, err := newSectorWeightBondModel(ctx, log, sectorWeightBond)
			if err != nil {
				continue
			}

			sectorWeightBondModels = append(sectorWeightBondModels, sectorWeightBondModel)
		}

		fundHoldingModel.Bonds = sectorWeightBondModels
		fundHoldingModel.Stocks = sectorWeightStockModels
	}

	return fundHoldingModel, nil
}

func newSectorWeightBondModel(ctx context.Context, log logger.ContextLog, sectorWeightBond *entities.SectorWeightBond) (*SectorWeightBondModel, error) {
	var sectorWeightBondModel = &SectorWeightBondModel{}

	if sectorWeightBond.MarketValPercent != "" {
		marketValPercent, err := sectorWeightBond.MarketValPercent.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightBond.MarketValPercent failed", "error", err, "MarketValPercent", sectorWeightBond.MarketValPercent)
			marketValPercent = 0
		}

		sectorWeightBondModel.MarketValPercent = marketValPercent
	}

	if sectorWeightBond.MarketValue != "" {
		marketValue, err := sectorWeightBond.MarketValue.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightBond.MarketValue failed", "error", err, "MarketValue", sectorWeightBond.MarketValue)
			marketValue = 0
		}

		sectorWeightBondModel.MarketValue = marketValue
	}

	if sectorWeightBond.Type != "" {
		sectorWeightBondModel.Type = sectorWeightBond.Type
	}

	sectorWeightBondModel.FaceAmount = sectorWeightBond.FaceAmount

	sectorWeightBondModel.Rate = sectorWeightBond.Rate

	return sectorWeightBondModel, nil
}

func newSectorWeightStockModel(ctx context.Context, log logger.ContextLog, sectorWeightStock *entities.SectorWeightStock) (*SectorWeightStockModel, error) {
	var sectorWeightStockModel = &SectorWeightStockModel{}

	if sectorWeightStock.Symbol != "" {
		sectorWeightStockModel.Symbol = sectorWeightStock.Symbol
	}

	if sectorWeightStock.MarketValPercent != "" {
		marketValPercent, err := sectorWeightStock.MarketValPercent.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightStock.MarketValPercent failed", "error", err, "MarketValPercent", sectorWeightStock.MarketValPercent)
			marketValPercent = 0
		}

		sectorWeightStockModel.MarketValPercent = marketValPercent
	}

	if sectorWeightStock.MarketValue != "" {
		marketValue, err := sectorWeightStock.MarketValue.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightStock.MarketValue failed", "error", err, "MarketValue", sectorWeightStock.MarketValue)
			marketValue = 0
		}

		sectorWeightStockModel.MarketValue = marketValue
	}

	if sectorWeightStock.Type != "" {
		sectorWeightStockModel.Type = sectorWeightStock.Type
	}

	sectorWeightStockModel.Shares = sectorWeightStock.Shares

	return sectorWeightStockModel, nil
}
