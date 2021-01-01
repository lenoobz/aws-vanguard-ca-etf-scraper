package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// IndividualFundModel is the representation of single vanguard fund model
type IndividualFundModel struct {
	ID            *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Schema        int                 `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive      bool                `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt     int64               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt    int64               `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Ticker        string              `json:"TICKER,omitempty" bson:"ticker,omitempty"`
	AssetCode     string              `json:"assetCode,omitempty" bson:"assetCode,omitempty"`
	Name          string              `json:"parentLongName,omitempty" bson:"name,omitempty"`
	Currency      string              `json:"currency,omitempty" bson:"currency,omitempty"`
	IssueTypeCode string              `json:"issueTypeCode,omitempty" bson:"issueTypeCode,omitempty"`
	PortID        string              `json:"portId,omitempty" bson:"portId,omitempty"`
	ProductType   string              `json:"productType,omitempty" bson:"productType,omitempty"`
	ManagementFee string              `json:"managementFee,omitempty" bson:"managementFee,omitempty"`
	MerValue      string              `json:"merValue,omitempty" bson:"merValue,omitempty"`
}
