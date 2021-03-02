package models

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardFundHoldingModel represents Vanguard fund holding model
type VanguardFundHoldingModel struct {
	ID         *primitive.ObjectID       `bson:"_id,omitempty"`
	IsActive   bool                      `bson:"isActive,omitempty"`
	CreatedAt  int64                     `bson:"createdAt,omitempty"`
	ModifiedAt int64                     `bson:"modifiedAt,omitempty"`
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

// NewHoldingModel create a fund holding model
func NewHoldingModel(ctx context.Context, log logger.ContextLog, e *entities.VanguardFundHolding) (*VanguardFundHoldingModel, error) {
	if e.AssetCode == "BOND" {
		return newBondHolding(ctx, log, e)
	}

	if e.AssetCode == "EQUITY" {
		return newEquityHolding(ctx, log, e)
	}

	if e.AssetCode == "BALANCED" {
		return newBalanceHolding(ctx, log, e)
	}

	return nil, nil
}

func newBondHolding(ctx context.Context, log logger.ContextLog, e *entities.VanguardFundHolding) (*VanguardFundHoldingModel, error) {
	var m = &VanguardFundHoldingModel{}

	if e.PortID != "" {
		m.PortID = e.PortID
	}

	if e.Ticker != "" {
		m.Ticker = e.Ticker
	}

	m.AssetCode = e.AssetCode

	if len(e.Bonds) == 1 {
		var bonds []*SectorWeightBondModel
		holding := e.Bonds[0]

		for _, v := range holding.SectorWeightBond {
			bond, err := newSectorWeightBondModel(ctx, log, v)
			if err != nil {
				continue
			}

			bonds = append(bonds, bond)
		}

		m.Bonds = bonds
	}

	return m, nil
}

func newEquityHolding(ctx context.Context, log logger.ContextLog, e *entities.VanguardFundHolding) (*VanguardFundHoldingModel, error) {
	var m = &VanguardFundHoldingModel{}

	if e.PortID != "" {
		m.PortID = e.PortID
	}

	if e.Ticker != "" {
		m.Ticker = e.Ticker
	}

	m.AssetCode = e.AssetCode

	if len(e.Equities) == 1 {
		var stocks []*SectorWeightStockModel
		holding := e.Equities[0]

		for _, v := range holding.SectorWeightStock {
			stock, err := newSectorWeightStockModel(ctx, log, v)
			if err != nil {
				continue
			}

			stocks = append(stocks, stock)
		}

		m.Stocks = stocks
	}

	return m, nil
}

func newBalanceHolding(ctx context.Context, log logger.ContextLog, e *entities.VanguardFundHolding) (*VanguardFundHoldingModel, error) {
	var m = &VanguardFundHoldingModel{}

	if e.PortID != "" {
		m.PortID = e.PortID
	}

	if e.Ticker != "" {
		m.Ticker = e.Ticker
	}

	m.AssetCode = e.AssetCode

	if len(e.Balances) == 1 {
		var stocks []*SectorWeightStockModel
		var bonds []*SectorWeightBondModel
		holding := e.Balances[0]

		for _, v := range holding.SectorWeightStock {
			stock, err := newSectorWeightStockModel(ctx, log, v)
			if err != nil {
				continue
			}

			stocks = append(stocks, stock)
		}

		for _, v := range holding.SectorWeightBond {
			bond, err := newSectorWeightBondModel(ctx, log, v)
			if err != nil {
				continue
			}

			bonds = append(bonds, bond)
		}

		m.Bonds = bonds
		m.Stocks = stocks
	}

	return m, nil
}

func newSectorWeightBondModel(ctx context.Context, log logger.ContextLog, e *entities.SectorWeightBond) (*SectorWeightBondModel, error) {
	var m = &SectorWeightBondModel{}

	if e.MarketValPercent != "" {
		v, err := e.MarketValPercent.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightBond.MarketValPercent failed", "err", err, "MarketValPercent", e.MarketValPercent)
			v = 0
		}

		m.MarketValPercent = v
	}

	if e.MarketValue != "" {
		v, err := e.MarketValue.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightBond.MarketValue failed", "err", err, "MarketValue", e.MarketValue)
			v = 0
		}

		m.MarketValue = v
	}

	if e.Type != "" {
		m.Type = e.Type
	}

	m.FaceAmount = e.FaceAmount

	m.Rate = e.Rate

	return m, nil
}

func newSectorWeightStockModel(ctx context.Context, log logger.ContextLog, e *entities.SectorWeightStock) (*SectorWeightStockModel, error) {
	var m = &SectorWeightStockModel{}

	if e.Symbol != "" {
		m.Symbol = e.Symbol
	}

	if e.MarketValPercent != "" {
		v, err := e.MarketValPercent.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightStock.MarketValPercent failed", "err", err, "MarketValPercent", e.MarketValPercent)
			v = 0
		}

		m.MarketValPercent = v
	}

	if e.MarketValue != "" {
		v, err := e.MarketValPercent.Float64()
		if err != nil {
			log.Warn(ctx, "parse SectorWeightStock.MarketValue failed", "err", err, "MarketValue", e.MarketValue)
			v = 0
		}

		m.MarketValue = v
	}

	if e.Type != "" {
		m.Type = e.Type
	}

	m.Shares = e.Shares

	return m, nil
}
