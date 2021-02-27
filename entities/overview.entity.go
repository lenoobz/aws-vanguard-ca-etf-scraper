package entities

// Overview represents Vanguard's fund overview
type Overview struct {
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

// FundCodesData represents fund's meta data such as ticker name, sedol code, isin number
type FundCodesData struct {
	Isin           string `json:"isin,omitempty"`
	Sedol          string `json:"sedol,omitempty"`
	ExchangeTicker string `json:"exchangeTicker,omitempty"`
}

// SectorWeighting represents fund's sector weight details
type SectorWeighting struct {
	FundPercent string `json:"fundPercent,omitempty"`
	LongName    string `json:"longName,omitempty"`
	SectorType  string `json:"sectorType,omitempty"`
}

// CountryExposure represents fund's country exposure details
type CountryExposure struct {
	CountryName     string `json:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty"`
}
