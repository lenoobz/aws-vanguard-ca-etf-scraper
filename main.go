package main

import (
	"fmt"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/repos"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/scrapers"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func main() {
	repo, err := repos.NewFundRepo(db)
	if err != nil {
		fmt.Println("error occurred when connect to database", err)
	}

	// we won't close database connection
	db = repo.DB

	// init service
	svc := services.NewFundService(repo)

	// init scrape
	jobs := scrapers.NewFundScraper(svc)

	if err := jobs.FundListCol.Visit(consts.FundListURL); err != nil {
		fmt.Println("error occurred when visit url", consts.FundListURL, err)
	}

	jobs.FundListCol.Wait()
	jobs.FundOverviewCol.Wait()
}
