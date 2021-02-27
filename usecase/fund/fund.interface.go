package fund

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Fund Repository Interface
///////////////////////////////////////////////////////////

// IFundReader interface
type IFundReader interface{}

// IFundWriter interface
type IFundWriter interface {
	InsertFund(ctx context.Context, fund *entities.Fund) error
}

// IFundRepo interface
type IFundRepo interface {
	IFundReader
	IFundWriter
}

///////////////////////////////////////////////////////////
// Fund Service Interface
///////////////////////////////////////////////////////////

// IFundService define business rule of fund
type IFundService interface {
	CreateFund(ctx context.Context, fund *entities.Fund) error
}
