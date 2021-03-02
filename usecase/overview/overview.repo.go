package overview

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Overview Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertOverview(context.Context, *entities.VanguardFundOverview) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
