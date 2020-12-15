package domains

// FundOverview is the representation of individual Vanguard fund overview
type FundOverview struct {
	PortID           string             `json:"portId,omitempty"`
	AssetClass       string             `json:"assetClass,omitempty"`
	Strategy         string             `json:"strategy,omitempty"`
	TotalAssets      string             `json:"totalAssets,omitempty"`
	DividendSchedule string             `json:"dividendSchedule,omitempty"`
	ShortName        string             `json:"shortName,omitempty"`
	Yield12Month     string             `json:"yield12Month,omitempty"`
	Price            float64            `json:"price,omitempty"`
	BaseCurrency     string             `json:"baseCurrency,omitempty"`
	ManagementFee    string             `json:"managementFee,omitempty"`
	MerValue         string             `json:"merValue,omitempty"`
	FundCodesData    *FundCodesData     `json:"fundCodesData,omitempty"`
	SectorWeighting  []*SectorWeighting `json:"sectorWeighting,omitempty"`
	CountryExposure  []*CountryExposure `json:"countryExposure,omitempty"`
}

// FundCodesData is the representation of meta data of the fund
type FundCodesData struct {
	Isin           string `json:"isin,omitempty"`
	Sedol          string `json:"sedol,omitempty"`
	ExchangeTicker string `json:"exchangeTicker,omitempty"`
}

// SectorWeighting is the representation of sector the fund invested
type SectorWeighting struct {
	FundPercent string `json:"fundPercent,omitempty"`
	LongName    string `json:"longName,omitempty"`
	SectorType  string `json:"sectorType,omitempty"`
}

// CountryExposure is the representation of country the fund exposed
type CountryExposure struct {
	CountryName     string `json:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty"`
}
