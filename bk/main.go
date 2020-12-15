package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type fundListResp struct {
	FundData map[string]fundData `json:"fundData,omitempty"`
}

type fundData struct {
	Ticker             string `json:"TICKER,omitempty"`
	AssetCode          string `json:"assetCode,omitempty"`
	ManagementTypeCode string `json:"managementTypeCode,omitempty"`
	ManagementFee      string `json:"managementFee,omitempty"`
	MerValue           string `json:"merValue,omitempty"`
	ShortName          string `json:"shortName,omitempty"`
	Currency           string `json:"currency,omitempty"`
}

type fundOverviewResp struct {
	PortID           string             `json:"portId,omitempty" bson:"portId,omitempty"`
	AssetClass       string             `json:"assetClass,omitempty" bson:"assetClass,omitempty"`
	Strategy         string             `json:"strategy,omitempty" bson:"strategy,omitempty"`
	TotalAssets      string             `json:"totalAssets,omitempty" bson:"totalAssets,omitempty"`
	DividendSchedule string             `json:"dividendSchedule,omitempty" bson:"dividendSchedule,omitempty"`
	ShortName        string             `json:"shortName,omitempty" bson:"shortName,omitempty"`
	Yield12Month     string             `json:"yield12Month,omitempty" bson:"yield12Month,omitempty"`
	Price            float64            `json:"price,omitempty" bson:"price,omitempty"`
	BaseCurrency     string             `json:"baseCurrency,omitempty" bson:"baseCurrency,omitempty"`
	ManagementFee    string             `json:"managementFee,omitempty" bson:"managementFee,omitempty"`
	MerValue         string             `json:"merValue,omitempty" bson:"merValue,omitempty"`
	FundCodesData    []*fundCodesData   `json:"fundCodesData,omitempty" bson:"fundCodesData,omitempty"`
	SectorWeighting  []*sectorWeighting `json:"sectorWeighting,omitempty" bson:"sectorWeighting,omitempty"`
	CountryExposure  []*countryExposure `json:"countryExposure,omitempty" bson:"countryExposure,omitempty"`
}

type fundCodesData struct {
	Isin           string `json:"isin,omitempty" bson:"isin,omitempty"`
	Sedol          string `json:"sedol,omitempty" bson:"sedol,omitempty"`
	ExchangeTicker string `json:"exchangeTicker,omitempty" bson:"exchangeTicker,omitempty"`
}

type sectorWeighting struct {
	FundPercent string `json:"fundPercent,omitempty" bson:"fundPercent,omitempty"`
	LongName    string `json:"longName,omitempty" bson:"longName,omitempty"`
	SectorType  string `json:"sectorType,omitempty" bson:"sectorType,omitempty"`
}

type countryExposure struct {
	CountryName     string `json:"countryName,omitempty" bson:"countryName,omitempty"`
	FundMktPercent  string `json:"fundMktPercent,omitempty" bson:"fundMktPercent,omitempty"`
	FundTnaPercent  string `json:"fundTnaPercent,omitempty" bson:"fundTnaPercent,omitempty"`
	HoldingStatCode string `json:"holdingStatCode,omitempty" bson:"holdingStatCode,omitempty"`
}

type fundHoldingResp struct {
	SectorWeightStock []*sectorWeightStock `json:"sectorWeightStock,omitempty"`
	SectorWeightBond  []*sectorWeightBond  `json:"sectorWeightBond,omitempty"`
}

type sectorWeightStock struct {
	Holding          string  `json:"holding,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
	Type             string  `json:"type,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
	MarketValue      float64 `json:"marketValue,omitempty"`
	Shares           float64 `json:"shares,omitempty"`
	Currency         string  `json:"currency,omitempty"`
}

type sectorWeightBond struct {
	Securities       string  `json:"securities,omitempty"`
	Type             string  `json:"type,omitempty"`
	Rate             float64 `json:"rate,omitempty"`
	MarketValPercent float64 `json:"marketValPercent,omitempty"`
}

var (
	con         *mongo.Database
	host        string
	dbname      string
	username    string
	password    string
	fundOverCol = "fund_overview"
)
var baseSearchURL = "https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-listview-data-en.json"

func getFundOverviewURL(portID string) string {
	return fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-overview-data-etf.json?vars=portId:%s,lang:en&path=[portId=%s][0]", portID, portID)
}

func getFundHoldingURL(portID, issueType, assetCode string) string {
	var URL string

	switch assetCode {
	case "BOND":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-bond.jsonp?vars=portId:%s,issueType:%s", portID, issueType)
		break
	case "EQUITY":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-equity.json?vars=portId:%s,issueType:%s", portID, issueType)
		break
	case "BALANCED":
		URL = fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-balanced.json?vars=portId:%s,issueType:%s", portID, issueType)
		break
	default:
		break
	}

	return URL
}

// func init() {
// 	// Read HOST environment variable
// 	host = os.Getenv("HOST")
// 	if host == "" {
// 		fmt.Println("error HOST is required")
// 	}

// 	// Read USERNAME environment variable
// 	username = os.Getenv("USERNAME")
// 	if username == "" {
// 		fmt.Println("error USERNAME is required")
// 	}

// 	// Read PASSWORD environment variable
// 	password = os.Getenv("PASSWORD")
// 	if password == "" {
// 		fmt.Println("error PASSWORD is required")
// 	}

// 	// Read DBNAME environment variable
// 	dbname = os.Getenv("DBNAME")
// 	if dbname == "" {
// 		fmt.Println("error DBNAME is required")
// 	}
// }

func init() {
	host = "lenoobetfdevcluster.jd7wd.mongodb.net"
	username = "lenoob_dev"
	password = "lenoob_dev"
	dbname = "etf_funds"
}

// newConnection creates new connection
func newConnection() (*mongo.Database, error) {
	// construct a connection string from mongo config object
	cxnString := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, host)

	// create mongo client by making new connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cxnString))
	if err != nil {
		return nil, err
	}

	// tell client what database to work with
	db := client.Database(dbname)

	// everything is looking good
	return db, nil
}

func scrapeFundList(fundOverviewCol, fundHoldingCol *colly.Collector, cb func(*colly.Collector, *colly.Collector, *fundListResp) error) {
	c := colly.NewCollector(
		// colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("api.vanguard.com"),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Handle response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing fund list response from Vanguard")
		rs := &fundListResp{}
		err := json.Unmarshal(r.Body, rs)
		if err != nil {
			fmt.Println("error occurred when parsing fund list response", err)
			return
		}

		if err = cb(fundOverviewCol, fundHoldingCol, rs); err != nil {
			fmt.Println("error occurred when processing fund list response", err)
			return
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	if err := c.Visit(baseSearchURL); err != nil {
		fmt.Println("error occurred when visit url", baseSearchURL, err)
	}
}

func scrapeFundOverview(cb func(*fundOverviewResp) error) *colly.Collector {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("api.vanguard.com"),
		colly.Async(true),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*vanguard.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Handle response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing fund overview response from Vanguard")
		rs := &fundOverviewResp{}
		err := json.Unmarshal(r.Body, rs)
		if err != nil {
			fmt.Println("error occurred when parsing fund overview response", err)
			return
		}

		if err = cb(rs); err != nil {
			fmt.Println("error occurred when processing fund overview response", err)
			return
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	return c
}

func scrapeFundHolding(cb func() error) *colly.Collector {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("api.vanguard.com"),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*vanguard.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Handle response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Parsing fund holding response from Vanguard")
		var rs []*fundHoldingResp
		err := json.Unmarshal(r.Body, &rs)
		if err != nil {
			fmt.Println("error occurred when parsing fund holding response", err)
			return
		}

		if err = cb(); err != nil {
			fmt.Println("error occurred when processing fund holding response", err)
			return
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	return c
}

func processFundList(fundOverviewCol *colly.Collector, fundHoldingCol *colly.Collector, fl *fundListResp) error {
	bondCounter := 0
	equityCounter := 0
	balancedCounter := 0
	for key, element := range fl.FundData {
		if element.Ticker != "" {
			fmt.Println("Key:", key, "=>", "Ticker:", element.Ticker, "=>", "Asset Code:", element.AssetCode)

			overviewURL := getFundOverviewURL(key)
			if err := fundOverviewCol.Visit(overviewURL); err != nil {
				fmt.Println("error occurred when visit url", overviewURL, err)
			}

			if element.AssetCode == "BOND" {
				bondCounter++
				continue
			}

			if element.AssetCode == "BALANCED" {
				balancedCounter++
				continue
			}

			if element.AssetCode == "EQUITY" {
				equityCounter++
				continue
			}
		}
	}

	fmt.Println("There are", bondCounter, "bond")
	fmt.Println("There are", equityCounter, "equity")
	fmt.Println("There are", balancedCounter, "balanced")
	return nil
}

func processFundOverview(fo *fundOverviewResp) error {
	// what collection we are going to use
	col := con.Collection(fundOverCol)

	// insert options
	insertOptions := options.InsertOne()

	_, err := col.InsertOne(context.Background(), fo, insertOptions)
	if err != nil {
		fmt.Println("error occurred when insert fund overview to db", err)
		return err
	}

	return nil
}

func processFundHolding() error {
	return nil
}

func main() {
	if con == nil {
		var err error
		con, err = newConnection()
		if err != nil {
			fmt.Println("error occurred when connect to database", err)
		}
	}

	fo := scrapeFundOverview(processFundOverview)
	fh := scrapeFundHolding(processFundHolding)
	scrapeFundList(fo, fh, processFundList)

	// Wait until threads are finished
	fo.Wait()
	fh.Wait()
}
