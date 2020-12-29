package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// IndividualFundModel is the representation of single vanguard fund model
type IndividualFundModel struct {
	ID         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Schema     int                 `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive   bool                `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt  int64               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt int64               `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Ticker     string              `json:"TICKER,omitempty" bson:"ticker,omitempty"`
	AssetCode  string              `json:"assetCode,omitempty" bson:"assetCode,omitempty"`
	Name       string              `json:"parentLongName,omitempty" bson:"name,omitempty"`
}
