package models

import (
	"context"
	"strconv"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HoldingModel represents Vanguard's fund holding details model
type HoldingModel struct {
	ID           *primitive.ObjectID       `json:"id,omitempty" bson:"_id,omitempty"`
	IsActive     bool                      `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt    int64                     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt   int64                     `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Schema       string                    `json:"schema,omitempty" bson:"schema,omitempty"`
	PortID       string                    `json:"portId,omitempty" bson:"portId,omitempty"`
	Ticker       string                    `json:"ticker,omitempty" bson:"ticker,omitempty"`
	AssetCode    string                    `json:"assetCode,omitempty" bson:"assetCode,omitempty"`
	BondHolding  []*SectorWeightBondModel  `json:"bondHolding,omitempty" bson:"bondHolding,omitempty"`
	StockHolding []*SectorWeightStockModel `json:"stockHolding,omitempty" bson:"stockHolding,omitempty"`
}

// SectorWeightBondModel struct
type SectorWeightBondModel struct {
	FaceAmount       float64 `json:"faceAmount,omitempty" bson:"faceAmount,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty" bson:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty" bson:"marketValue,omitempty"`
	Rate             float64 `json:"rate,omitempty" bson:"rate,omitempty"`
	Type             string  `json:"type,omitempty" bson:"type,omitempty"`
}

// SectorWeightStockModel struct
type SectorWeightStockModel struct {
	MarketValPercent float64 `json:"marketValPercent,omitempty" bson:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty" bson:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty" bson:"shares,omitempty"`
	Symbol           string  `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Type             string  `json:"type,omitempty" bson:"type,omitempty"`
}

// NewHoldingModel create a fund holding model
func NewHoldingModel(ctx context.Context, log logger.IAppLogger, h *entities.Holding) (*HoldingModel, error) {
	if h.AssetCode == "BOND" {
		return newBondHolding(ctx, log, h)
	}

	if h.AssetCode == "EQUITY" {
		return newEquityHolding(ctx, log, h)
	}

	if h.AssetCode == "BALANCED" {
		return newBalanceHolding(ctx, log, h)
	}

	return nil, nil
}

func newBondHolding(ctx context.Context, log logger.IAppLogger, h *entities.Holding) (*HoldingModel, error) {
	var holding = &HoldingModel{}

	if h.PortID != "" {
		holding.PortID = h.PortID
	}

	if h.Ticker != "" {
		holding.Ticker = h.Ticker
	}

	holding.AssetCode = h.AssetCode

	if len(h.BondHolding) == 1 {
		var bonds []*SectorWeightBondModel
		bondHolding := h.BondHolding[0]

		for _, v := range bondHolding.SectorWeightBond {
			bond, err := newBondSectorWeightBondModel(ctx, log, v)

			if err != nil {
				log.Error(ctx, "map BondHoldingSectorWeightBond failed", "err", err)
				continue
			}

			bonds = append(bonds, bond)
		}

		holding.BondHolding = bonds
	}

	return holding, nil
}

func newBondSectorWeightBondModel(ctx context.Context, log logger.IAppLogger, bond *entities.BondHoldingSectorWeightBond) (*SectorWeightBondModel, error) {
	var bondModel = &SectorWeightBondModel{}

	bondModel.FaceAmount = bond.FaceAmount

	if bond.MarketValPercent != "" {
		marketValPercent, err := strconv.ParseFloat(bond.MarketValPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse BondHoldingSectorWeightBond.MarketValPercent failed", "err", err, "MarketValPercent", bond.MarketValPercent)
			marketValPercent = 0
		}

		bondModel.MarketValPercent = marketValPercent
	}

	bondModel.MarketValue = bond.MarketValue

	bondModel.Rate = bond.Rate

	bondModel.Type = "BOND"

	return bondModel, nil
}

func newEquityHolding(ctx context.Context, log logger.IAppLogger, h *entities.Holding) (*HoldingModel, error) {
	var holding = &HoldingModel{}

	if h.PortID != "" {
		holding.PortID = h.PortID
	}

	if h.Ticker != "" {
		holding.Ticker = h.Ticker
	}

	holding.AssetCode = h.AssetCode

	if len(h.EquityHolding) == 1 {
		var stocks []*SectorWeightStockModel
		stockHolding := h.EquityHolding[0]

		for _, v := range stockHolding.SectorWeightStock {
			stock, err := newEquityHoldingSectorWeightStock(ctx, log, v)

			if err != nil {
				log.Error(ctx, "map EquityHoldingSectorWeightStock failed", "err", err)
				continue
			}

			stocks = append(stocks, stock)
		}

		holding.StockHolding = stocks
	}

	return holding, nil
}

func newEquityHoldingSectorWeightStock(ctx context.Context, log logger.IAppLogger, equity *entities.EquityHoldingSectorWeightStock) (*SectorWeightStockModel, error) {
	var equityModel = &SectorWeightStockModel{}

	if equity.Symbol != "" {
		equityModel.Symbol = equity.Symbol
	}

	if equity.MarketValPercent != "" {
		marketValPercent, err := strconv.ParseFloat(equity.MarketValPercent, 64)

		if err != nil {
			log.Warn(ctx, "parse EquityHoldingSectorWeightStock.MarketValPercent failed", "err", err, "MarketValPercent", equity.MarketValPercent)
			marketValPercent = 0
		}

		equityModel.MarketValPercent = marketValPercent
	}

	if equity.MarketValue != "" {
		marketValue, err := strconv.ParseFloat(equity.MarketValue, 64)

		if err != nil {
			log.Warn(ctx, "parse EquityHoldingSectorWeightStock.MarketValue failed", "err", err, "MarketValue", equity.MarketValue)
			marketValue = 0
		}

		equityModel.MarketValue = marketValue
	}

	equityModel.Shares = equity.Shares

	equityModel.Type = "STOCK"

	return equityModel, nil
}

func newBalanceHolding(ctx context.Context, log logger.IAppLogger, h *entities.Holding) (*HoldingModel, error) {
	var holding = &HoldingModel{}

	if h.PortID != "" {
		holding.PortID = h.PortID
	}

	if h.Ticker != "" {
		holding.Ticker = h.Ticker
	}

	holding.AssetCode = h.AssetCode

	if len(h.BalancedHolding) == 1 {
		var stocks []*SectorWeightStockModel
		var bonds []*SectorWeightBondModel
		balancedHolding := h.BalancedHolding[0]

		for _, v := range balancedHolding.SectorWeightStock {
			stock, err := newBalancedHoldingSectorWeightStock(ctx, log, v)

			if err != nil {
				log.Error(ctx, "map BalancedHoldingSectorWeightStock failed", "err", err)
				continue
			}

			stocks = append(stocks, stock)
		}

		for _, v := range balancedHolding.SectorWeightBond {
			bond, err := newBalancedHoldingSectorWeightBond(ctx, log, v)

			if err != nil {
				log.Error(ctx, "map BalancedHoldingSectorWeightBond failed", "err", err)
				continue
			}

			bonds = append(bonds, bond)
		}

		holding.BondHolding = bonds
		holding.StockHolding = stocks
	}

	return holding, nil
}

func newBalancedHoldingSectorWeightBond(ctx context.Context, log logger.IAppLogger, bond *entities.BalancedHoldingSectorWeightBond) (*SectorWeightBondModel, error) {
	var bondModel = &SectorWeightBondModel{}

	bondModel.MarketValPercent = bond.MarketValPercent

	bondModel.Rate = bond.Rate

	if bond.Type != "" {
		bondModel.Type = bond.Type
	}

	return bondModel, nil
}

func newBalancedHoldingSectorWeightStock(ctx context.Context, log logger.IAppLogger, equity *entities.BalancedHoldingSectorWeightStock) (*SectorWeightStockModel, error) {
	var equityModel = &SectorWeightStockModel{}

	if equity.Symbol != "" {
		equityModel.Symbol = equity.Symbol
	}

	equityModel.MarketValPercent = equity.MarketValPercent

	equityModel.MarketValue = equity.MarketValue

	equityModel.Shares = equity.Shares

	equityModel.Type = equity.Type

	return equityModel, nil
}
