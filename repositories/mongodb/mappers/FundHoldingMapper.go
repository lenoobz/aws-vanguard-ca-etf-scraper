package mappers

import (
	"fmt"
	"strconv"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/domains"
	"github.com/hthl85/aws-vanguard-ca-etf-scraper/repositories/mongodb/models"
)

// MapFundHoldingDomain converts fund holding domain to model
func MapFundHoldingDomain(fund *domains.FundHolding) (*models.FundHoldingModel, error) {
	var fundHolding = &models.FundHoldingModel{}

	if fund.PortID != "" {
		fundHolding.PortID = fund.PortID
	}

	if fund.Ticker != "" {
		fundHolding.Ticker = fund.Ticker
	}

	if fund.AssetCode == "BOND" {
		fundHolding.AssetCode = fund.AssetCode

		if len(fund.BondHolding) == 1 {
			var bonds []*models.SectorWeightBondModel
			bondHolding := fund.BondHolding[0]

			for _, v := range bondHolding.SectorWeightBond {
				bond, err := mapBondHoldingSectorWeightBond(v)

				if err != nil {
					fmt.Println("error occurred when mapping BondHoldingSectorWeightBond field", err)
					continue
				}

				bonds = append(bonds, bond)
			}

			fundHolding.BondHolding = bonds
		}

		return fundHolding, nil
	}

	if fund.AssetCode == "EQUITY" {
		fundHolding.AssetCode = fund.AssetCode

		if len(fund.EquityHolding) == 1 {
			var stocks []*models.SectorWeightStockModel
			stockHolding := fund.EquityHolding[0]

			for _, v := range stockHolding.SectorWeightStock {
				stock, err := mapEquityHoldingSectorWeightStock(v)

				if err != nil {
					fmt.Println("error occurred when mapping EquityHoldingSectorWeightStock field", err)
					continue
				}

				stocks = append(stocks, stock)
			}

			fundHolding.StockHolding = stocks
		}

		return fundHolding, nil
	}

	if fund.AssetCode == "BALANCED" {
		fundHolding.AssetCode = fund.AssetCode

		if len(fund.BalancedHolding) == 1 {
			var stocks []*models.SectorWeightStockModel
			var bonds []*models.SectorWeightBondModel
			balancedHolding := fund.BalancedHolding[0]

			for _, v := range balancedHolding.SectorWeightStock {
				stock, err := mapBalancedHoldingSectorWeightStock(v)

				if err != nil {
					fmt.Println("error occurred when mapping BalancedHoldingSectorWeightStock field", err)
					continue
				}

				stocks = append(stocks, stock)
			}

			for _, v := range balancedHolding.SectorWeightBond {
				bond, err := mapBalancedHoldingSectorWeightBond(v)

				if err != nil {
					fmt.Println("error occurred when mapping BalancedHoldingSectorWeightBond field", err)
					continue
				}

				bonds = append(bonds, bond)
			}

			fundHolding.BondHolding = bonds
			fundHolding.StockHolding = stocks
		}

		return fundHolding, nil
	}

	return fundHolding, nil
}

func mapBondHoldingSectorWeightBond(bond *domains.BondHoldingSectorWeightBond) (*models.SectorWeightBondModel, error) {
	var bondModel = &models.SectorWeightBondModel{}

	// if bond.Cusip != "" {
	// 	bondModel.Cusip = bond.Cusip
	// }

	// if bond.Isin != "" {
	// 	bondModel.Isin = bond.Isin
	// }

	// if bond.MaturityDate != "" {
	// 	bondModel.MaturityDate = bond.MaturityDate
	// }

	// if bond.MaturityDateNumber != "" {
	// 	bondModel.MaturityDateNumber = bond.MaturityDateNumber
	// }

	// if bond.Securities != "" {
	// 	bondModel.Securities = bond.Securities
	// }

	// if bond.Sedol != "" {
	// 	bondModel.Sedol = bond.Sedol
	// }

	bondModel.FaceAmount = bond.FaceAmount

	if bond.MarketValPercent != "" {
		marketValPercent, err := strconv.ParseFloat(bond.MarketValPercent, 64)

		if err != nil {
			fmt.Println("error occurred when parsing marketValPercent field", err)
			marketValPercent = 0
		}

		bondModel.MarketValPercent = marketValPercent
	}

	bondModel.MarketValue = bond.MarketValue

	bondModel.Rate = bond.Rate

	bondModel.Type = "BOND"

	return bondModel, nil
}

func mapEquityHoldingSectorWeightStock(equity *domains.EquityHoldingSectorWeightStock) (*models.SectorWeightStockModel, error) {
	var equityModel = &models.SectorWeightStockModel{}

	// if equity.Currency != "" {
	// 	equityModel.Currency = equity.Currency
	// }

	// if equity.Country != "" {
	// 	equityModel.Country = equity.Country
	// }

	// if equity.Cusip != "" {
	// 	equityModel.Cusip = equity.Cusip
	// }

	// if equity.EquityExchangeCode != "" {
	// 	equityModel.EquityExchangeCode = equity.EquityExchangeCode
	// }

	// if equity.Holding != "" {
	// 	equityModel.Holding = equity.Holding
	// }

	// if equity.Isin != "" {
	// 	equityModel.Holding = equity.Isin
	// }

	// if equity.Sector != "" {
	// 	equityModel.Holding = equity.Sector
	// }

	// if equity.Sedol != "" {
	// 	equityModel.Holding = equity.Sedol
	// }

	if equity.Symbol != "" {
		equityModel.Symbol = equity.Symbol
	}

	if equity.MarketValPercent != "" {
		marketValPercent, err := strconv.ParseFloat(equity.MarketValPercent, 64)

		if err != nil {
			fmt.Println("error occurred when parsing marketValPercent field", err)
			marketValPercent = 0
		}

		equityModel.MarketValPercent = marketValPercent
	}

	if equity.MarketValue != "" {
		marketValue, err := strconv.ParseFloat(equity.MarketValue, 64)

		if err != nil {
			fmt.Println("error occurred when parsing marketValue field", err)
			marketValue = 0
		}

		equityModel.MarketValue = marketValue
	}

	equityModel.Shares = equity.Shares

	equityModel.Type = "STOCK"

	return equityModel, nil
}

func mapBalancedHoldingSectorWeightBond(bond *domains.BalancedHoldingSectorWeightBond) (*models.SectorWeightBondModel, error) {
	var bondModel = &models.SectorWeightBondModel{}

	// if bond.MaturityDate != "" {
	// 	bondModel.MaturityDate = bond.MaturityDate
	// }

	// if bond.MaturityDateNumber != "" {
	// 	bondModel.MaturityDateNumber = bond.MaturityDateNumber
	// }

	// if bond.Securities != "" {
	// 	bondModel.Securities = bond.Securities
	// }

	bondModel.MarketValPercent = bond.MarketValPercent

	bondModel.Rate = bond.Rate

	if bond.Type != "" {
		bondModel.Type = bond.Type
	}

	return bondModel, nil
}

func mapBalancedHoldingSectorWeightStock(equity *domains.BalancedHoldingSectorWeightStock) (*models.SectorWeightStockModel, error) {
	var equityModel = &models.SectorWeightStockModel{}

	// if equity.Currency != "" {
	// 	equityModel.Currency = equity.Currency
	// }

	// if equity.Holding != "" {
	// 	equityModel.Holding = equity.Holding
	// }

	if equity.Symbol != "" {
		equityModel.Symbol = equity.Symbol
	}

	equityModel.MarketValPercent = equity.MarketValPercent

	equityModel.MarketValue = equity.MarketValue

	equityModel.Shares = equity.Shares

	equityModel.Type = equity.Type

	return equityModel, nil
}
