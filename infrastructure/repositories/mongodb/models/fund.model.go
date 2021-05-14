package models

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardFundModel represents a Vanguard fund model
type VanguardFundModel struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"`
	IsActive      bool                `bson:"isActive,omitempty"`
	CreatedAt     int64               `bson:"createdAt,omitempty"`
	ModifiedAt    int64               `bson:"modifiedAt,omitempty"`
	Schema        string              `bson:"schema,omitempty"`
	Ticker        string              `bson:"ticker,omitempty"`
	AssetCode     string              `bson:"assetCode,omitempty"`
	Name          string              `bson:"name,omitempty"`
	Currency      string              `bson:"currency,omitempty"`
	IssueType     string              `bson:"issueType,omitempty"`
	PortID        string              `bson:"portId,omitempty"`
	ProductType   string              `bson:"productType,omitempty"`
	ManagementFee string              `bson:"managementFee,omitempty"`
	MerFee        string              `bson:"merFee,omitempty"`
}

// NewVanguardFundModel create Vanguard fund model
func NewVanguardFundModel(vanguardFund *entities.VanguardFund) (*VanguardFundModel, error) {
	var fundModel = &VanguardFundModel{}

	if vanguardFund.Ticker != "" {
		fundModel.Ticker = vanguardFund.Ticker
	}

	if vanguardFund.Name != "" {
		fundModel.Name = vanguardFund.Name
	}

	if vanguardFund.AssetCode != "" {
		fundModel.AssetCode = vanguardFund.AssetCode
	}

	if vanguardFund.Currency != "" {
		fundModel.Currency = vanguardFund.Currency
	}

	if vanguardFund.IssueType != "" {
		fundModel.IssueType = vanguardFund.IssueType
	}

	if vanguardFund.PortID != "" {
		fundModel.PortID = vanguardFund.PortID
	}

	if vanguardFund.ProductType != "" {
		fundModel.ProductType = vanguardFund.ProductType
	}

	if vanguardFund.ManagementFee != "" {
		fundModel.ManagementFee = vanguardFund.ManagementFee
	}

	if vanguardFund.MerFee != "" {
		fundModel.MerFee = vanguardFund.MerFee
	}

	return fundModel, nil
}
