package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FundHoldingModel is the representation of individual Vanguard fund overview model
type FundHoldingModel struct {
	ID           *primitive.ObjectID       `json:"id,omitempty" bson:"_id,omitempty"`
	Schema       int                       `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive     bool                      `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt    int64                     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt   int64                     `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
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
