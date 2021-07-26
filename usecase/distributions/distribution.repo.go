package distributions

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
	InsertFundDistribution(ctx context.Context, fundDistribution *entities.FundDistribution) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
