package holding

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Holding Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertHolding(context.Context, *entities.VanguardFundHolding) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
