package entities

import "encoding/json"

// VanguardFundHolding represents Vanguard fund holding entity
type VanguardFundHolding struct {
	PortID    string             `json:"portId,omitempty"`
	Ticker    string             `json:"ticker,omitempty"`
	AssetCode string             `json:"assetCode,omitempty"`
	Bonds     []*BondHolding     `json:"bondHolding,omitempty"`
	Equities  []*EquityHolding   `json:"equityHolding,omitempty"`
	Balances  []*BalancedHolding `json:"balancedHolding,omitempty"`
}

// BondHolding represents fund bond holding details
type BondHolding struct {
	SectorWeightBonds []*SectorWeightBond `json:"sectorWeightBond,omitempty"`
}

// EquityHolding represents fund equity holding details
type EquityHolding struct {
	SectorWeightStocks []*SectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// BalancedHolding represents fund balance holding details
type BalancedHolding struct {
	SectorWeightBonds  []*SectorWeightBond  `json:"sectorWeightBond,omitempty"`
	SectorWeightStocks []*SectorWeightStock `json:"sectorWeightStock,omitempty"`
}

///////////////////////////////////////////////////////////
// Holding detail inner structs
///////////////////////////////////////////////////////////

// SectorWeightBond struct
type SectorWeightBond struct {
	MarketValPercent json.Number `json:"marketValPercent,omitempty"`
	MarketValue      json.Number `json:"marketValue,omitempty"`
	FaceAmount       float64     `json:"faceAmount,omitempty"`
	Rate             float64     `json:"rate,omitempty"`
	Type             string      `json:"type,omitempty"`
}

// SectorWeightStock struct
type SectorWeightStock struct {
	MarketValPercent json.Number `json:"marketValPercent,omitempty"`
	MarketValue      json.Number `json:"marketValue,omitempty"`
	Shares           float64     `json:"shares,omitempty"`
	Symbol           string      `json:"symbol,omitempty"`
	Type             string      `json:"type,omitempty"`
}
