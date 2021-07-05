package models

import (
	"context"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundModel struct
type FundModel struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt     int64               `bson:"createdAt,omitempty"`
	ModifiedAt    int64               `bson:"modifiedAt,omitempty"`
	Enabled       bool                `bson:"enabled"`
	Deleted       bool                `bson:"deleted"`
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

// NewFundModel create Vanguard fund model
func NewFundModel(ctx context.Context, log logger.ContextLog, vanguardFund *entities.Fund, schemaVersion string) (*FundModel, error) {
	var fundModel = &FundModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if vanguardFund.Ticker != "" {
		fundModel.Ticker = ticker.GenYahooTickerFromVanguardTicker(vanguardFund.Ticker)
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
