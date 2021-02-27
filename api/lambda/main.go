package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/repositories/mongodb/repos"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/scraper"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/fund"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/overview"
)

func main() {
	appConf := config.AppConf

	// create new logger
	logger, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer logger.Close()

	// create new repository
	repo, err := repos.NewFundMongo(nil, logger, &appConf.Mongo)
	if err != nil {
		log.Fatal("create fund mongo repo failed")
	}
	defer repo.Close()

	// create new service
	fundService := fund.NewFundService(repo, logger)
	holdingService := holding.NewHoldingService(repo, logger)
	overviewService := overview.NewOverviewService(repo, logger)

	// create new scraper jobs
	jobs := scraper.NewFundScraper(fundService, holdingService, overviewService, logger)
	jobs.StartJobs()

	lambda.Start(lambdaHandler)
}

func lambdaHandler() {
	log.Println("lambda handler is called")
}
