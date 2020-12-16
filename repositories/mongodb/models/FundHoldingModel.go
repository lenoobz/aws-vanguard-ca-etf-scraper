package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FundHoldingModel is the representation of individual Vanguard fund overview model
type FundHoldingModel struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Schema    int                 `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive  bool                `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt int64               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
