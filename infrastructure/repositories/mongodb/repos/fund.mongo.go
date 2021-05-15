package repos

import (
	"context"
	"fmt"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/entities"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/repositories/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FundMongo struct
type FundMongo struct {
	db     *mongo.Database
	client *mongo.Client
	log    logger.ContextLog
	conf   *config.MongoConfig
}

// NewFundMongo creates new fund mongo repo
func NewFundMongo(db *mongo.Database, log logger.ContextLog, conf *config.MongoConfig) (*FundMongo, error) {
	if db != nil {
		return &FundMongo{
			db:   db,
			log:  log,
			conf: conf,
		}, nil
	}

	// set context with timeout from the config
	// create new context for the query
	ctx, cancel := createContext(context.Background(), conf.TimeoutMS)
	defer cancel()

	// set mongo client options
	clientOptions := options.Client()

	// set min pool size
	if conf.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(conf.MinPoolSize)
	}

	// set max pool size
	if conf.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(conf.MaxPoolSize)
	}

	// set max idle time ms
	if conf.MaxIdleTimeMS > 0 {
		clientOptions.SetMaxConnIdleTime(time.Duration(conf.MaxIdleTimeMS) * time.Millisecond)
	}

	// construct a connection string from mongo config object
	cxnString := fmt.Sprintf("mongodb+srv://%s:%s@%s", conf.Username, conf.Password, conf.Host)

	// create mongo client by making new connection
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(cxnString))
	if err != nil {
		return nil, err
	}

	return &FundMongo{
		db:     client.Database(conf.Dbname),
		client: client,
		log:    log,
		conf:   conf,
	}, nil
}

// Close disconnect from database
func (r *FundMongo) Close() {
	ctx := context.Background()
	r.log.Info(ctx, "close mongo client")

	if r.client == nil {
		return
	}

	if err := r.client.Disconnect(ctx); err != nil {
		r.log.Error(ctx, "disconnect mongo failed", "error", err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Implement interface
///////////////////////////////////////////////////////////////////////////////

// InsertFund inserts new fund
func (r *FundMongo) InsertFund(ctx context.Context, fund *entities.VanguardFund) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	fundModel, err := models.NewFundModel(fund)
	if err != nil {
		r.log.Error(ctx, "create model failed", "error", err)
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames[consts.VANGUARD_FUNDS_COL]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
		return fmt.Errorf("cannot find collection name")
	}
	col := r.db.Collection(colname)

	fundModel.IsActive = true
	fundModel.Schema = r.conf.SchemaVersion
	fundModel.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: fundModel.Ticker,
	}}

	update := bson.D{
		{
			Key:   "$set",
			Value: fundModel,
		},
		{
			Key: "$setOnInsert",
			Value: bson.D{{
				Key:   "createdAt",
				Value: time.Now().UTC().Unix(),
			}},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.log.Error(ctx, "update one failed", "error", err)
		return err
	}

	return nil
}

// InsertFundOverview inserts fund overview
func (r *FundMongo) InsertFundOverview(ctx context.Context, fundOverview *entities.VanguardFundOverview) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	fundOverviewModel, err := models.NewFundOverviewModel(ctx, r.log, fundOverview)
	if err != nil {
		r.log.Error(ctx, "create model failed", "error", err)
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames[consts.VANGUARD_OVERVIEW_COL]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	fundOverviewModel.IsActive = true
	fundOverviewModel.Schema = r.conf.SchemaVersion
	fundOverviewModel.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: fundOverviewModel.Ticker,
	}}

	update := bson.D{
		{
			Key:   "$set",
			Value: fundOverviewModel,
		},
		{
			Key: "$setOnInsert",
			Value: bson.D{{
				Key:   "createdAt",
				Value: time.Now().UTC().Unix(),
			}},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.log.Error(ctx, "update one failed", "error", err)
		return err
	}

	return nil
}

// InsertFundHolding inserts fund holding
func (r *FundMongo) InsertFundHolding(ctx context.Context, fundHolding *entities.VanguardFundHolding) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	fundHoldingModel, err := models.NewFundHoldingModel(ctx, r.log, fundHolding)
	if err != nil {
		r.log.Error(ctx, "create model failed", "error", err)
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames[consts.VANGUARD_HOLDING_COL]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	fundHoldingModel.IsActive = true
	fundHoldingModel.Schema = r.conf.SchemaVersion
	fundHoldingModel.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: fundHoldingModel.Ticker,
	}}

	update := bson.D{
		{
			Key:   "$set",
			Value: fundHoldingModel,
		},
		{
			Key: "$setOnInsert",
			Value: bson.D{{
				Key:   "createdAt",
				Value: time.Now().UTC().Unix(),
			}},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.log.Error(ctx, "update one failed", "error", err)
		return err
	}

	return nil
}

// InsertFundDistribution inserts fund distribution
func (r *FundMongo) InsertFundDistribution(ctx context.Context, fundDistribution *entities.VanguardFundDistribution) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	fundDistributionModel, err := models.NewFundDistributionModel(ctx, r.log, fundDistribution)
	if err != nil {
		r.log.Error(ctx, "create model failed", "error", err)
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames[consts.VANGUARD_DISTRIBUTION_COL]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	fundDistributionModel.IsActive = true
	fundDistributionModel.Schema = r.conf.SchemaVersion
	fundDistributionModel.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "portId",
		Value: fundDistributionModel.PortID,
	}}

	update := bson.D{
		{
			Key:   "$set",
			Value: fundDistributionModel,
		},
		{
			Key: "$setOnInsert",
			Value: bson.D{{
				Key:   "createdAt",
				Value: time.Now().UTC().Unix(),
			}},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.log.Error(ctx, "update one failed", "error", err)
		return err
	}

	return nil
}

///////////////////////////////////////////////////////////
// Implement helper function
///////////////////////////////////////////////////////////

// createContext create a new context with timeout
func createContext(ctx context.Context, t uint64) (context.Context, context.CancelFunc) {
	timeout := time.Duration(t) * time.Millisecond
	return context.WithTimeout(ctx, timeout*time.Millisecond)
}
