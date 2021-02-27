package overview

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
)

///////////////////////////////////////////////////////////
// Overview Repository Interface
///////////////////////////////////////////////////////////

// IOverviewReader interface
type IOverviewReader interface{}

// IOverviewWriter interface
type IOverviewWriter interface {
	InsertOverview(ctx context.Context, overview *entities.Overview) error
}

// IOverviewRepo interface
type IOverviewRepo interface {
	IOverviewReader
	IOverviewWriter
}

///////////////////////////////////////////////////////////
// Overivew Service Interface
///////////////////////////////////////////////////////////

// IOverviewService define business rule of overview
type IOverviewService interface {
	CreateOverview(ctx context.Context, overview *entities.Overview) error
}
