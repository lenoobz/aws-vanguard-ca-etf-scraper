package models

import (
	"context"
	"time"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/entities"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundDistributionModel struct
type FundDistributionModel struct {
	ID                    *primitive.ObjectID         `bson:"_id,omitempty"`
	CreatedAt             int64                       `bson:"createdAt,omitempty"`
	ModifiedAt            int64                       `bson:"modifiedAt,omitempty"`
	Enabled               bool                        `bson:"enabled"`
	Deleted               bool                        `bson:"deleted"`
	Schema                string                      `bson:"schema,omitempty"`
	PortID                string                      `bson:"portId,omitempty"`
	Ticker                string                      `bson:"ticker,omitempty"`
	DistributionHistories []*DistributionHistoryModel `bson:"distributionHistories,omitempty"`
}

// FundDistribution struct
type DistributionHistoryModel struct {
	Type               string  `bson:"type,omitempty"`
	DistributionAmount float64 `bson:"distributionAmount,omitempty"`
	ExDividendDate     string  `bson:"exDividendDate,omitempty"`
	RecordDate         string  `bson:"recordDate,omitempty"`
	PayableDate        string  `bson:"payableDate,omitempty"`
	DistDesc           string  `bson:"distDesc,omitempty"`
	DistCode           string  `bson:"distCode,omitempty"`
}

// NewFundDistributionModel create a fund distribution model
func NewFundDistributionModel(ctx context.Context, log logger.ContextLog, fundDistribution *entities.FundDistribution, schemaVersion string) (*FundDistributionModel, error) {
	distributionDetails := fundDistribution.DistributionDetails
	var fundDistributionModels = &FundDistributionModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
	}

	if distributionDetails.PortID != "" {
		fundDistributionModels.PortID = distributionDetails.PortID
	}

	if distributionDetails.Ticker != "" {
		fundDistributionModels.Ticker = ticker.GenYahooTickerFromVanguardTicker(distributionDetails.Ticker)
	}

	// map distribution histories model
	var distributionHistoryModels []*DistributionHistoryModel
	for _, distributionHistory := range distributionDetails.DistributionHistories {
		distributionHistoryModel, err := newDistributionHistoryModel(ctx, log, distributionHistory)

		if err != nil {
			return nil, err
		}

		distributionHistoryModels = append(distributionHistoryModels, distributionHistoryModel)
	}
	fundDistributionModels.DistributionHistories = distributionHistoryModels

	return fundDistributionModels, nil
}

func newDistributionHistoryModel(ctx context.Context, log logger.ContextLog, distributionHistory *entities.DistributionHistory) (*DistributionHistoryModel, error) {
	return &DistributionHistoryModel{
		Type:               distributionHistory.Type,
		DistributionAmount: distributionHistory.DistributionAmount,
		ExDividendDate:     distributionHistory.ExDividendDate,
		RecordDate:         distributionHistory.RecordDate,
		PayableDate:        distributionHistory.PayableDate,
		DistDesc:           distributionHistory.DistDesc,
		DistCode:           distributionHistory.DistCode,
	}, nil
}
