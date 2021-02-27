package holding

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Holding Repository Interface
///////////////////////////////////////////////////////////

// IHoldingReader interface
type IHoldingReader interface{}

// IHoldingWriter interface
type IHoldingWriter interface {
	InsertHolding(ctx context.Context, holding *entities.Holding) error
}

// IHoldingRepo interface
type IHoldingRepo interface {
	IHoldingReader
	IHoldingWriter
}

///////////////////////////////////////////////////////////
// Holding Service Interface
///////////////////////////////////////////////////////////

// IHoldingService define business rule of holding
type IHoldingService interface {
	CreateHolding(ctx context.Context, holding *entities.Holding) error
}
