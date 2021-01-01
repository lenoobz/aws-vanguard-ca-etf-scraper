package domains

// VanguardFunds is the representation of individual Vanguard fund collections
type VanguardFunds struct {
	IndividualFunds map[string]IndividualFund `json:"fundData,omitempty"`
}

// IndividualFund is the representation of individual Vanguard fund
type IndividualFund struct {
	Ticker        string `json:"TICKER,omitempty"`
	AssetCode     string `json:"assetCode,omitempty"`
	Name          string `json:"parentLongName,omitempty"`
	Currency      string `json:"currency,omitempty"`
	IssueTypeCode string `json:"issueTypeCode,omitempty"`
	PortID        string `json:"portId,omitempty"`
	ProductType   string `json:"productType,omitempty"`
	ManagementFee string `json:"managementFee,omitempty"`
	MerValue      string `json:"merValue,omitempty"`
}
