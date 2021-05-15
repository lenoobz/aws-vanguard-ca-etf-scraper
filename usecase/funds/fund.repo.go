package funds

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Fund Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertFund(ctx context.Context, fund *entities.VanguardFund) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}