package entities

// VanguardFundOverview represents Vanguard fund overview entity
type VanguardFundOverview struct {
	PortID           string              `json:"portId,omitempty"`
	AssetClass       string              `json:"assetClass,omitempty"`
	Strategy         string              `json:"strategy,omitempty"`
	TotalAssets      string              `json:"totalAssets,omitempty"`
	DividendSchedule string              `json:"dividendSchedule,omitempty"`
	ShortName        string              `json:"shortName,omitempty"`
	Yield12Month     string              `json:"yield12Month,omitempty"`
	Price            float64             `json:"price,omitempty"`
	BaseCurrency     string              `json:"baseCurrency,omitempty"`
	ManagementFee    string              `json:"managementFee,omitempty"`
	MerFee           string              `json:"merValue,omitempty"`
	FundCode         *FundCode           `json:"fundCodesData,omitempty"`
	Sectors          []*SectorBreakdown  `json:"sectorWeighting,omitempty"`
	Countries        []*CountryBreakdown `json:"countryExposure,omitempty"`
}

// FundCode represents fund meta data such as ticker name, sedol code, isin number
type FundCode struct {
	Isin           string `json:"isin,omitempty"`
	Sedol          string `json:"sedol,omitempty"`
	ExchangeTicker string `json:"exchangeTicker,omitempty"`
}

// SectorBreakdown represents fund sector breakdown
type SectorBreakdown struct {
	FundPercent string `json:"fundPercent,omitempty"`
	SectorName  string `json:"longName,omitempty"`
}

// CountryBreakdown represents fund country breakdown
type CountryBreakdown struct {
	CountryName     string `json:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty"`
}
