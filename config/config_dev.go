// +build dev

package config

// AppConf constants
var AppConf = AppConfig{
	Mongo: MongoConfig{
		TimeoutMS:     360000,
		MinPoolSize:   5,
		MaxPoolSize:   10,
		MaxIdleTimeMS: 360000,
		Host:          "lenoobetfdevcluster.jd7wd.mongodb.net",
		Username:      "lenoob_dev",
		Password:      "lenoob_dev",
		Dbname:        "etf_funds_dev",
		SchemaVersion: "1",
		Colnames: map[string]string{
			"fund":     "vanguard_fund_lists",
			"overview": "vanguard_fund_overview",
			"holding":  "vanguard_fund_holding",
		},
	},
}

// Collection names
const (
	VANGUARD_FUNDS_COL    = "fund"
	VANGUARD_HOLDING_COL  = "holding"
	VANGUARD_OVERVIEW_COL = "overview"
)
