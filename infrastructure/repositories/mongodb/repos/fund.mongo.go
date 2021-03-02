package repos

import (
	"context"
	"fmt"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
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
func NewFundMongo(db *mongo.Database, l logger.ContextLog, conf *config.MongoConfig) (*FundMongo, error) {
	if db != nil {
		return &FundMongo{
			db:   db,
			log:  l,
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
		log:    l,
		conf:   conf,
	}, nil
}

// Close disconnect from database
func (r *FundMongo) Close() {
	ctx := context.Background()
	r.log.Info(ctx, "close fund mongo client")

	if r.client == nil {
		return
	}

	if err := r.client.Disconnect(ctx); err != nil {
		r.log.Error(ctx, "disconnect mongo failed", err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Implement interface
///////////////////////////////////////////////////////////////////////////////

// InsertFund inserts new fund
func (r *FundMongo) InsertFund(ctx context.Context, e *entities.VanguardFund) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	m, err := models.NewVanguardFundModel(e)
	if err != nil {
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames["fund"]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	m.IsActive = true
	m.Schema = r.conf.SchemaVersion
	m.CreatedAt = time.Now().UTC().Unix()
	m.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: m.Ticker,
	}}

	update := bson.D{{
		Key:   "$set",
		Value: m,
	}}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

// InsertOverview inserts fund overview
func (r *FundMongo) InsertOverview(ctx context.Context, e *entities.VanguardFundOverview) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	m, err := models.NewOverviewModel(ctx, r.log, e)
	if err != nil {
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames["overview"]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	m.IsActive = true
	m.Schema = r.conf.SchemaVersion
	m.CreatedAt = time.Now().UTC().Unix()
	m.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: m.Ticker,
	}}

	update := bson.D{{
		Key:   "$set",
		Value: m,
	}}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

// InsertHolding inserts fund holding
func (r *FundMongo) InsertHolding(ctx context.Context, e *entities.VanguardFundHolding) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	m, err := models.NewHoldingModel(ctx, r.log, e)
	if err != nil {
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames["holding"]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	m.IsActive = true
	m.Schema = r.conf.SchemaVersion
	m.CreatedAt = time.Now().UTC().Unix()
	m.ModifiedAt = time.Now().UTC().Unix()

	filter := bson.D{{
		Key:   "ticker",
		Value: m.Ticker,
	}}

	update := bson.D{{
		Key:   "$set",
		Value: m,
	}}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
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
