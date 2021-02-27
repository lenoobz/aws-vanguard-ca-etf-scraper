package models

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundModel represents Vanguard's individual fund model
type FundModel struct {
	ID            *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	IsActive      bool                `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt     int64               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt    int64               `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Schema        string              `json:"schema,omitempty" bson:"schema,omitempty"`
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

// NewFundModel create a fund model
func NewFundModel(f *entities.Fund) (*FundModel, error) {
	var fund = &FundModel{}

	if f.Ticker != "" {
		fund.Ticker = f.Ticker
	}

	if f.Name != "" {
		fund.Name = f.Name
	}

	if f.AssetCode != "" {
		fund.AssetCode = f.AssetCode
	}

	if f.Currency != "" {
		fund.Currency = f.Currency
	}

	if f.IssueTypeCode != "" {
		fund.IssueTypeCode = f.IssueTypeCode
	}

	if f.PortID != "" {
		fund.PortID = f.PortID
	}

	if f.ProductType != "" {
		fund.ProductType = f.ProductType
	}

	if f.ManagementFee != "" {
		fund.ManagementFee = f.ManagementFee
	}

	if f.MerValue != "" {
		fund.MerValue = f.MerValue
	}

	return fund, nil
}
