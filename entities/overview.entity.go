package entities

// VanguardFundOverview struct
type VanguardFundOverview struct {
	PortID           string              `json:"portId,omitempty"`
	AssetClass       string              `json:"assetClass,omitempty"`
	Strategy         string              `json:"strategy,omitempty"`
	TotalAssets      string              `json:"totalAssets,omitempty"`
	DividendSchedule string              `json:"dividendSchedule,omitempty"`
	Name             string              `json:"name,omitempty"`
	ShortName        string              `json:"shortName,omitempty"`
	Yield12Month     string              `json:"yield12Month,omitempty"`
	Price            float64             `json:"price,omitempty"`
	BaseCurrency     string              `json:"baseCurrency,omitempty"`
	ManagementFee    string              `json:"managementFee,omitempty"`
	MerFee           string              `json:"merValue,omitempty"`
	DistYield        string              `json:"distYield,omitempty"`
	DistAmount       string              `json:"incomeDistributionAmount,omitempty"`
	AllocationStock  float64             `json:"allocationStock,omitempty"`
	AllocationBond   float64             `json:"allocationBond,omitempty"`
	AllocationCash   float64             `json:"allocationCash,omitempty"`
	FundCode         *FundCode           `json:"fundCodesData,omitempty"`
	Sectors          []*SectorBreakdown  `json:"sectorWeighting,omitempty"`
	Countries        []*CountryBreakdown `json:"countryExposure,omitempty"`
	Dividends        []*DividendHistory  `json:"distHistory,omitempty"`
}

// FundCode struct
type FundCode struct {
	Isin           string `json:"isin,omitempty"`
	Sedol          string `json:"sedol,omitempty"`
	ExchangeTicker string `json:"exchangeTicker,omitempty"`
}

// SectorBreakdown struct
type SectorBreakdown struct {
	BenchmarkPercent string `json:"benchmarkPercent,omitempty"`
	FundPercent      string `json:"fundPercent,omitempty"`
	SectorName       string `json:"longName,omitempty"`
}

// CountryBreakdown struct
type CountryBreakdown struct {
	CountryName     string `json:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty"`
}

// DividendHistory struct
type DividendHistory struct {
	AsOfDate     string `json:"asOfDate,omitempty"`
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}
