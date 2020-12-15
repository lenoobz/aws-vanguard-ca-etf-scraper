package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FundOverviewModel is the representation of individual Vanguard fund overview model
type FundOverviewModel struct {
	ID               *primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Schema           int                     `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive         bool                    `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt        int64                   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	PortID           string                  `json:"portId,omitempty" bson:"portId,omitempty"`
	AssetClass       string                  `json:"assetClass,omitempty" bson:"assetClass,omitempty"`
	Strategy         string                  `json:"strategy,omitempty" bson:"strategy,omitempty"`
	DividendSchedule string                  `json:"dividendSchedule,omitempty" bson:"dividendSchedule,omitempty"`
	ShortName        string                  `json:"shortName,omitempty" bson:"shortName,omitempty"`
	Currency         string                  `json:"baseCurrency,omitempty" bson:"baseCurrency,omitempty"`
	Isin             string                  `json:"isin,omitempty" bson:"isin,omitempty"`
	Sedol            string                  `json:"sedol,omitempty" bson:"sedol,omitempty"`
	Ticker           string                  `json:"ticker,omitempty" bson:"ticker,omitempty"`
	TotalAssets      float64                 `json:"totalAssets,omitempty" bson:"totalAssets,omitempty"`
	Yield12Month     float64                 `json:"yield12Month,omitempty" bson:"yield12Month,omitempty"`
	Price            float64                 `json:"price,omitempty" bson:"price,omitempty"`
	ManagementFee    float64                 `json:"managementFee,omitempty" bson:"managementFee,omitempty"`
	MerValue         float64                 `json:"merValue,omitempty" bson:"merValue,omitempty"`
	SectorWeighting  []*SectorWeightingModel `json:"sectorWeighting,omitempty" bson:"sectorWeighting,omitempty"`
	CountryExposure  []*CountryExposureModel `json:"countryExposure,omitempty" bson:"countryExposure,omitempty"`
}

// SectorWeightingModel is the representation of sector the fund invested
type SectorWeightingModel struct {
	FundPercent float64 `json:"fundPercent,omitempty" bson:"fundPercent,omitempty"`
	LongName    string  `json:"longName,omitempty" bson:"longName,omitempty"`
	SectorType  string  `json:"sectorType,omitempty" bson:"sectorType,omitempty"`
}

// CountryExposureModel is the representation of country the fund exposed
type CountryExposureModel struct {
	CountryName     string  `json:"countryName,omitempty" bson:"countryName,omitempty"`
	HoldingStatCode string  `json:"holdingStatCode,omitempty" bson:"holdingStatCode,omitempty"`
	FundMktPercent  float64 `json:"fundMktPercent,omitempty" bson:"fundMktPercent,omitempty"`
	FundTnaPercent  float64 `json:"fundTnaPercent,omitempty" bson:"fundTnaPercent,omitempty"`
}
