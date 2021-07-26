package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/config"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/infrastructure/scraper"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/usecase/distributions"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/usecase/funds"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/lenoobz/aws-vanguard-ca-etf-scraper/usecase/overview"
)

func main() {
	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	repo, err := repos.NewFundMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create fund mongo repo failed")
	}
	defer repo.Close()

	// create new service
	fundService := funds.NewService(repo, zap)
	fundHoldingService := holding.NewService(repo, zap)
	fundOverviewService := overview.NewService(repo, zap)
	fundDistributionService := distributions.NewService(repo, zap)

	// create new scraper jobs
	jobs := scraper.NewFundScraper(fundService, fundHoldingService, fundOverviewService, fundDistributionService, zap)
	jobs.ScrapeAllVanguardFundsDetails()

	lambda.Start(lambdaHandler)
}

func lambdaHandler() {
	log.Println("lambda handler is called")
}
