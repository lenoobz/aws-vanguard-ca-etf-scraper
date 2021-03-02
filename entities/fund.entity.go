package entities

// VanguardFund represents a Vanguard fund entity
type VanguardFund struct {
	Ticker        string `json:"TICKER,omitempty"`
	AssetCode     string `json:"assetCode,omitempty"`
	Name          string `json:"parentLongName,omitempty"`
	Currency      string `json:"currency,omitempty"`
	IssueType     string `json:"issueTypeCode,omitempty"`
	PortID        string `json:"portId,omitempty"`
	ProductType   string `json:"productType,omitempty"`
	ManagementFee string `json:"managementFee,omitempty"`
	MerFee        string `json:"merValue,omitempty"`
}
