package domains

// VanguardFunds is the representation of individual Vanguard fund collections
type VanguardFunds struct {
	IndividualFunds map[string]IndividualFund `json:"fundData,omitempty"`
}

// IndividualFund is the representation of individual Vanguard fund
type IndividualFund struct {
	Ticker    string `json:"TICKER,omitempty"`
	AssetCode string `json:"assetCode,omitempty"`
	Name      string `json:"parentLongName,omitempty"`
}
