package overview

import (
	"context"

	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Overview Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertFundOverview(ctx context.Context, fundOverview *entities.FundOverview) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
