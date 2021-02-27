package entities

// Holding represents Vanguard's fund holding details
type Holding struct {
	PortID          string             `json:"portId,omitempty"`
	Ticker          string             `json:"ticker,omitempty"`
	AssetCode       string             `json:"assetCode,omitempty"`
	BondHolding     []*BondHolding     `json:"bondHolding,omitempty"`
	EquityHolding   []*EquityHolding   `json:"equityHolding,omitempty"`
	BalancedHolding []*BalancedHolding `json:"balancedHolding,omitempty"`
}

///////////////////////////////////////////////////////////
// BondHolding struct and its dependent structs
///////////////////////////////////////////////////////////

// BondHolding represents bond holding details of a fund
type BondHolding struct {
	SectorWeightBond []*BondHoldingSectorWeightBond `json:"sectorWeightBond,omitempty"`
}

// BondHoldingSectorWeightBond struct
type BondHoldingSectorWeightBond struct {
	FaceAmount       float64 `json:"faceAmount,omitempty"`
	MarketValPercent string  `json:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty"`
	Rate             float64 `json:"rate,omitempty"`
}

///////////////////////////////////////////////////////////
// EquityHolding struct and its dependent structs
///////////////////////////////////////////////////////////

// EquityHolding represents equity holding details of a fund
type EquityHolding struct {
	SectorWeightStock []*EquityHoldingSectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// EquityHoldingSectorWeightStock struct
type EquityHoldingSectorWeightStock struct {
	MarketValPercent string  `json:"marketValPercent,omitempty"`
	MarketValue      string  `json:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
}

///////////////////////////////////////////////////////////
// BalancedHolding struct and its dependent structs
///////////////////////////////////////////////////////////

// BalancedHolding represents balance holding details of a fund
type BalancedHolding struct {
	SectorWeightBond  []*BalancedHoldingSectorWeightBond  `json:"sectorWeightBond,omitempty"`
	SectorWeightStock []*BalancedHoldingSectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// BalancedHoldingSectorWeightBond struct
type BalancedHoldingSectorWeightBond struct {
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
	Rate             float64 `json:"rate,omitempty"`
	Type             string  `json:"type,omitempty"`
}

// BalancedHoldingSectorWeightStock struct
type BalancedHoldingSectorWeightStock struct {
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
	Type             string  `json:"type,omitempty"`
}
