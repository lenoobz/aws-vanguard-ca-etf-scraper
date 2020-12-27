package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/mappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// FundRepo struct
type FundRepo struct {
	DB *mongo.Database
}

// NewFundRepo creates new fund mongo repo
func NewFundRepo(db *mongo.Database) (*FundRepo, error) {
	if db != nil {
		return &FundRepo{
			DB: db,
		}, nil
	}

	// set context with timeout from the config
	timeout := time.Duration(consts.TimeoutMS) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// set mongo client options
	clientOptions := options.Client()

	// set min pool size
	if consts.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(consts.MinPoolSize)
	}

	// set max pool size
	if consts.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(consts.MaxPoolSize)
	}

	// set max idle time ms
	if consts.MaxIdleTimeMS > 0 {
		clientOptions.SetMaxConnIdleTime(time.Duration(consts.MaxIdleTimeMS) * time.Millisecond)
	}

	// construct a connection string from mongo config object
	cxnString := fmt.Sprintf("mongodb+srv://%s:%s@%s", consts.Username, consts.Password, consts.Host)

	// create mongo client by making new connection
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(cxnString))
	if err != nil {
		return nil, err
	}

	// test our connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &FundRepo{
		DB: client.Database(consts.Dbname),
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// Implement interface
///////////////////////////////////////////////////////////////////////////////

func (repo *FundRepo) isFundOverviewExisted(ticker string) (bool, error) {
	// create new context for the query
	ctx, cancel := createContext()
	defer cancel()

	// what collection we are going to use
	col := repo.DB.Collection(consts.CollectionFundOverview)

	filter := bson.D{{
		Key:   "ticker",
		Value: ticker,
	}}

	var fo *domains.FundOverview
	if err := col.FindOne(ctx, filter).Decode(fo); err != nil {
		return false, err
	}

	return true, nil
}

// InsertFundOverview inserts new fund overview
func (repo *FundRepo) InsertFundOverview(fo *domains.FundOverview) error {
	fundOverview, err := mappers.MapFundOverviewDomain(fo)
	if err != nil {
		return err
	}

	// create new context for the query
	ctx, cancel := createContext()
	defer cancel()

	// what collection we are going to use
	col := repo.DB.Collection(consts.CollectionFundOverview)

	fundOverview.IsActive = true
	fundOverview.Schema = consts.SchemaVersion
	fundOverview.CreatedAt = time.Now().UTC().Unix()
	fundOverview.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: fundOverview.Ticker,
	}}

	update := bson.D{{
		Key:   "$set",
		Value: fundOverview,
	}}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

// InsertFundHolding inserts new fund holding
func (repo *FundRepo) InsertFundHolding(fh *domains.FundHolding) error {
	fundHolding, err := mappers.MapFundHoldingDomain(fh)
	if err != nil {
		return err
	}

	// create new context for the query
	ctx, cancel := createContext()
	defer cancel()

	// what collection we are going to use
	col := repo.DB.Collection(consts.CollectionFundHolding)

	fundHolding.IsActive = true
	fundHolding.Schema = consts.SchemaVersion
	fundHolding.CreatedAt = time.Now().UTC().Unix()
	fundHolding.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: fundHolding.Ticker,
	}}

	update := bson.D{{
		Key:   "$set",
		Value: fundHolding,
	}}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
// Private helper functions
///////////////////////////////////////////////////////////////////////////////

// createContext create a new context with timeout
func createContext() (context.Context, context.CancelFunc) {
	// set context with timeout from the config
	timeout := time.Duration(consts.TimeoutMS) * time.Millisecond
	return context.WithTimeout(context.Background(), timeout*time.Millisecond)
}
