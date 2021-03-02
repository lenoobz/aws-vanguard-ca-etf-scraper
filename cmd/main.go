package main

import (
	"log"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/config"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/repositories/mongodb/repos"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/infrastructure/scraper"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/fund"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/holding"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/usecase/overview"
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
	fs := fund.NewService(repo, zap)
	hs := holding.NewService(repo, zap)
	os := overview.NewService(repo, zap)

	// create new scraper jobs
	jobs := scraper.NewFundScraper(fs, hs, os, zap)
	jobs.StartJobs()
}
