# Scrape Vanguard Canada ETF funds

The lambda function to scrape `Vanguard Canada ETF` data and extract interested data such as `Fund List`, `Fund Overview`, and `Fund Holding`.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Overview](#overview)
  - [Technical Summary](#technical-summary)
  - [Vanguard Endpoints](#vanguard-endpoints)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Build lambda function](#build-lambda-function)
  - [Build cmd](#build-cmd)
  - [Clean up](#clean-up)
- [How To](#how-to)
  - [Add new build environment](#add-new-build-environment)
- [Contributing](#contributing)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

#### Technical Summary

This lambda function is the starting point of a `state machine` which is used to scrape data from Vanguard Canada data.

The `state machine` was configurated in the `povi-infrastructure` project and was scheduled to run `Every Monday`. When the `state machine` is executed, this lambda function will be triggered.

The lambda function will scrape `Fund List`, `Fund Overview`, and `Fund Holding` data from the Vanguard Canada endpoints listed bellow.

Scraped data then will be parsed and stored in three mongo collections `vanguard_fund_lists`, `vanguard_fund_overview`, and `vanguard_fund_holding`.

_NOTE:_ Those collections are intended to use for raw data only.

#### Vanguard Endpoints

_NOTE: The `{portId}` value can be found from the `Fund List` data, and use `F` for {issueType}._

- `Fund List` endpoint (https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-listview-data-en.json)
- `Fund OverView` endpoint (https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-overview-data-etf.json?vars=portId:{portId},lang:en&path=[portId={portId}][0])
- `Fund Holding` have three separated endpoints for: `BOND`, `EQUITY`, and `BALANCED`.
  - `BOND holding` endpoint (https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-bond.json?vars=portId:{portId},issueType:{issueType})
  - `EQUITY holding` endpoint (https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-equity.json?vars=portId:{portId},issueType:{issueType})
  - `BALANCED holding` endpoint (https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-holding-details-balanced.json?vars=portId:{portId},issueType:{issueType})

## Project Structure

This project follows the clean architecture with minor variation

```
├── api
│   └── lambda
├── cmd
├── config
├── entities
├── infrastructure
│   ├── logger
│   ├── repositories
│   │   └── mongodb
│   │       ├── models
│   │       └── repos
│   └── scraper
├── usecase
│   ├── fund
│   ├── holding
│   ├── logger
│   └── overview
└── utils
    └── corid
```

## Usage

To build the project for different environment set env variable `LIBRARY_ENV`. Default environment is `dev`. Check [`How To`](#how-to) section to learn how to add new build env.

#### Build lambda function

Bellow command will build a `lambda function` to deploy to `AWS`. The output binary file is in `./bin/api/lambda/main

```bash
# Build lambda function
make build
```

#### Build cmd

Bellow command is to build a `command line` to run `locally`. The output binary file is in `./bin/cmd/main

```bash
# Build cmd
make build-cmd
```

#### Clean up

Bellow command is to clean up the build

```bash
# Clean
make clean
```

## How To

### Add new build environment

- Make a copy of `config_dev.go` and change the suffix to new environment name. For example, adding `staging` the new file should be named `config_staging.go`.
- Change the build tag (the first line) to new environment name. For example, changing `// +build dev` to `// +build staging`
- Finally, update the variable values in new config file accordingly.
- To build the project for the newly added environment don't forget to set env variable `LIBRARY_ENV`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

This project is not an opensource project. Please contact the owner for permission.
