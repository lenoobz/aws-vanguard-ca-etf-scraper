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
	SectorWeightBond []*SectorWeightBond `json:"sectorWeightBond,omitempty"`
}

// EquityHolding represents fund equity holding details
type EquityHolding struct {
	SectorWeightStock []*SectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// BalancedHolding represents fund balance holding details
type BalancedHolding struct {
	SectorWeightBond  []*SectorWeightBond  `json:"sectorWeightBond,omitempty"`
	SectorWeightStock []*SectorWeightStock `json:"sectorWeightStock,omitempty"`
}

///////////////////////////////////////////////////////////
// Holding detail inner structs
///////////////////////////////////////////////////////////

// SectorWeightBond struct
type SectorWeightBond struct {
	MarketValPercent json.Number `json:"marketValPercent,omitempty"`
	MarketValue      json.Number `json:"marketValue,omitempty"`
	Rate             float64     `json:"rate,omitempty"`
	FaceAmount       float64     `json:"faceAmount,omitempty"`
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
