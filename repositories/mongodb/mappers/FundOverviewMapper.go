package mappers

import (
	"fmt"
	"strconv"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/models"
)

// MapFundOverviewDomain converts fund overview domain to model
func MapFundOverviewDomain(fund *domains.FundOverview) (*models.FundOverviewModel, error) {
	var fundOverview = &models.FundOverviewModel{}

	if fund.PortID != "" {
		fundOverview.PortID = fund.PortID
	}

	if fund.AssetClass != "" {
		fundOverview.AssetClass = fund.AssetClass
	}

	if fund.Strategy != "" {
		fundOverview.Strategy = fund.Strategy
	}

	if fund.DividendSchedule != "" {
		fundOverview.DividendSchedule = fund.DividendSchedule
	}

	if fund.ShortName != "" {
		fundOverview.ShortName = fund.ShortName
	}

	if fund.BaseCurrency != "" {
		fundOverview.Currency = fund.BaseCurrency
	}

	if fund.FundCodesData != nil {
		if fund.FundCodesData.Isin != "" {
			fundOverview.Isin = fund.FundCodesData.Isin
		}

		if fund.FundCodesData.Sedol != "" {
			fundOverview.Sedol = fund.FundCodesData.Sedol
		}

		if fund.FundCodesData.ExchangeTicker != "" {
			fundOverview.Ticker = fund.FundCodesData.ExchangeTicker
		}
	}

	if fund.TotalAssets != "" {
		totalAssets, err := strconv.ParseFloat(fund.TotalAssets, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing totalAssets field", err)
			totalAssets = 0
		}

		fundOverview.TotalAssets = totalAssets
	}

	if fund.Yield12Month != "" {
		yield12Month, err := strconv.ParseFloat(fund.Yield12Month, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing yield12Month field", err)
			yield12Month = 0
		}

		fundOverview.Yield12Month = yield12Month
	}

	fundOverview.Price = fund.Price

	if fund.ManagementFee != "" {
		managementFee, err := strconv.ParseFloat(fund.ManagementFee, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing managementFee field", err)
			managementFee = 0
		}

		fundOverview.ManagementFee = managementFee
	}

	if fund.MerValue != "" {
		merValue, err := strconv.ParseFloat(fund.MerValue, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing merValue field", err)
			merValue = 0
		}

		fundOverview.MerValue = merValue
	}

	var sectors []*models.SectorWeightingModel
	for _, v := range fund.SectorWeighting {
		sector, err := mapSectorWeighting(v)

		if err != nil {
			return nil, err
		}

		sectors = append(sectors, sector)
	}
	fundOverview.SectorWeighting = sectors

	var exposures []*models.CountryExposureModel
	for _, v := range fund.CountryExposure {
		exposure, err := mapCountryExposure(v)

		if err != nil {
			return nil, err
		}

		if exposure != nil {
			exposures = append(exposures, exposure)
		}
	}
	fundOverview.CountryExposure = exposures

	return fundOverview, nil
}

func mapSectorWeighting(sector *domains.SectorWeighting) (*models.SectorWeightingModel, error) {
	var sectorWeighting = &models.SectorWeightingModel{}

	if sector.LongName != "" {
		sectorWeighting.LongName = sector.LongName
	}

	if sector.SectorType != "" {
		sectorWeighting.SectorType = sector.SectorType
	}

	if sector.FundPercent != "" {
		fundPercent, err := strconv.ParseFloat(sector.FundPercent, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing fundPercent field", err)
			fundPercent = 0
		}

		sectorWeighting.FundPercent = fundPercent
	}

	return sectorWeighting, nil
}

func mapCountryExposure(exposure *domains.CountryExposure) (*models.CountryExposureModel, error) {
	var countryExposure = &models.CountryExposureModel{}

	if exposure.CountryName != "" {
		countryExposure.CountryName = exposure.CountryName
	}

	if exposure.HoldingStatCode != "" {
		countryExposure.HoldingStatCode = exposure.HoldingStatCode
	}

	if exposure.FundMktPercent != "" {
		mktPercent, err := strconv.ParseFloat(exposure.FundMktPercent, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing mktPercent field", err)
			mktPercent = 0
		}

		if mktPercent == 0 {
			return nil, nil
		}

		countryExposure.FundMktPercent = mktPercent
	}

	if exposure.FundTnaPercent != "" {
		tnaPercent, err := strconv.ParseFloat(exposure.FundTnaPercent, 64)

		if err != nil {
			fmt.Println("[warning] - error occurred when parsing tnaPercent field", err)
			tnaPercent = 0
		}

		countryExposure.FundTnaPercent = tnaPercent
	}

	return countryExposure, nil
}
