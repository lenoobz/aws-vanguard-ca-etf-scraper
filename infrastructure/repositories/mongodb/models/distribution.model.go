package models

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/ticker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardFundDistributionModel struct
type VanguardFundDistributionModel struct {
	ID                    *primitive.ObjectID         `bson:"_id,omitempty"`
	IsActive              bool                        `bson:"isActive,omitempty"`
	CreatedAt             int64                       `bson:"createdAt,omitempty"`
	ModifiedAt            int64                       `bson:"modifiedAt,omitempty"`
	Schema                string                      `bson:"schema,omitempty"`
	PortID                string                      `bson:"portId,omitempty"`
	Ticker                string                      `bson:"ticker,omitempty"`
	DistributionHistories []*DistributionHistoryModel `bson:"distributionHistories,omitempty"`
}

// FundDistribution struct
type DistributionHistoryModel struct {
	Type               string  `bson:"portId,omitempty"`
	DistributionAmount float64 `bson:"distributionAmount,omitempty"`
	ExDividendDate     string  `bson:"exDividendDate,omitempty"`
	RecordDate         string  `bson:"recordDate,omitempty"`
	PayableDate        string  `bson:"payableDate,omitempty"`
	DistDesc           string  `bson:"distDesc,omitempty"`
	DistCode           string  `bson:"distCode,omitempty"`
}

// NewFundDistributionModel create a fund distribution model
func NewFundDistributionModel(ctx context.Context, log logger.ContextLog, fundDistribution *entities.VanguardFundDistribution) (*VanguardFundDistributionModel, error) {
	distributionDetails := fundDistribution.DistributionDetails
	var fundDistributionModels = &VanguardFundDistributionModel{}

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
