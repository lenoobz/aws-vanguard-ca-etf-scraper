package funds

import (
	"context"

	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Fund Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertFund(ctx context.Context, fund *entities.Fund) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
