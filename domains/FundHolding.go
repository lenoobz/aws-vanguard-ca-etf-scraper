package domains

// FundHolding struct
type FundHolding struct {
	PortID          string             `json:"portId,omitempty"`
	Ticker          string             `json:"ticker,omitempty"`
	AssetCode       string             `json:"assetCode,omitempty"`
	BondHolding     []*BondHolding     `json:"bondHolding,omitempty"`
	EquityHolding   []*EquityHolding   `json:"equityHolding,omitempty"`
	BalancedHolding []*BalancedHolding `json:"balancedHolding,omitempty"`
}

// BondHolding struct
type BondHolding struct {
	SectorWeightBond []*BondHoldingSectorWeightBond `json:"sectorWeightBond,omitempty"`
}

// BondHoldingSectorWeightBond struct
type BondHoldingSectorWeightBond struct {
	CurrencySymbol     string  `json:"currencySymbol,omitempty"`
	Cusip              string  `json:"cusip,omitempty"`
	FaceAmount         float64 `json:"faceAmount,omitempty"`
	Isin               string  `json:"isin,omitempty"`
	MarketValPercent   string  `json:"marketValPercent,omitempty"`
	MarketValue        float64 `json:"marketValue,omitempty"`
	MaturityDate       string  `json:"maturitydate,omitempty"`
	MaturityDateNumber string  `json:"maturitydatenumber,omitempty"`
	Rate               float64 `json:"rate,omitempty"`
	Securities         string  `json:"securities,omitempty"`
	Sedol              string  `json:"sedol,omitempty"`
}

// EquityHolding struct
type EquityHolding struct {
	SectorWeightStock []*EquityHoldingSectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// EquityHoldingSectorWeightStock struct
type EquityHoldingSectorWeightStock struct {
	Currency           string  `json:"currency,omitempty"`
	Country            string  `json:"country,omitempty"`
	Cusip              string  `json:"cusip,omitempty"`
	EquityExchangeCode string  `json:"equityExchangeCode,omitempty"`
	Holding            string  `json:"holding,omitempty"`
	Isin               string  `json:"isin,omitempty"`
	MarketValPercent   string  `json:"marketValPercent,omitempty"`
	MarketValue        string  `json:"marketValue,omitempty"`
	Sector             string  `json:"sector,omitempty"`
	Sedol              string  `json:"sedol,omitempty"`
	Shares             float64 `json:"shares,omitempty"`
	Symbol             string  `json:"symbol,omitempty"`
}

// BalancedHolding struct
type BalancedHolding struct {
	SectorWeightBond  []*BalancedHoldingSectorWeightBond  `json:"sectorWeightBond,omitempty"`
	SectorWeightStock []*BalancedHoldingSectorWeightStock `json:"sectorWeightStock,omitempty"`
}

// BalancedHoldingSectorWeightBond struct
type BalancedHoldingSectorWeightBond struct {
	MarketValPercent   float64 `json:"marketValPercent,omitempty"`
	MaturityDate       string  `json:"maturitydate,omitempty"`
	MaturityDateNumber string  `json:"maturitydatenumber,omitempty"`
	Rate               float64 `json:"rate,omitempty"`
	Securities         string  `json:"securities,omitempty"`
	Type               string  `json:"type,omitempty"`
}

// BalancedHoldingSectorWeightStock struct
type BalancedHoldingSectorWeightStock struct {
	Currency         string  `json:"currency,omitempty"`
	Holding          string  `json:"holding,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
	Type             string  `json:"type,omitempty"`
}
