package entities

// VanguardFundDistribution struct
type VanguardFundDistribution struct {
	DistributionDetails struct {
		PortID                string                 `json:"portId,omitempty"`
		Ticker                string                 `json:"ticker,omitempty"`
		DistributionHistories []*DistributionHistory `json:"fundDistributionList,omitempty"`
	} `json:"distributions,omitempty"`
}

// DistributionHistory struct
type DistributionHistory struct {
	Type               string  `json:"type,omitempty"`
	DistributionAmount float64 `json:"distributionAmount,omitempty"`
	ExDividendDate     string  `json:"exDividendDate,omitempty"`
	RecordDate         string  `json:"recordDate,omitempty"`
	PayableDate        string  `json:"payableDate,omitempty"`
	DistDesc           string  `json:"distDesc,omitempty"`
	DistCode           string  `json:"distCode,omitempty"`
}
