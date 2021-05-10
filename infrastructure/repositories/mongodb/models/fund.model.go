package models

import (
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/ticker"
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
func NewVanguardFundModel(e *entities.VanguardFund) (*VanguardFundModel, error) {
	var m = &VanguardFundModel{}

	if e.Ticker != "" {
		m.Ticker = ticker.GetYahooTicker(e.Ticker)
	}

	if e.Name != "" {
		m.Name = e.Name
	}

	if e.AssetCode != "" {
		m.AssetCode = e.AssetCode
	}

	if e.Currency != "" {
		m.Currency = e.Currency
	}

	if e.IssueType != "" {
		m.IssueType = e.IssueType
	}

	if e.PortID != "" {
		m.PortID = e.PortID
	}

	if e.ProductType != "" {
		m.ProductType = e.ProductType
	}

	if e.ManagementFee != "" {
		m.ManagementFee = e.ManagementFee
	}

	if e.MerFee != "" {
		m.MerFee = e.MerFee
	}

	return m, nil
}
