// +build local

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
		Dbname:        "povi_local",
		SchemaVersion: "1",
		Colnames: map[string]string{
			"vanguard_fund_list":         "vanguard_fund_list",
			"vanguard_fund_overview":     "vanguard_fund_overview",
			"vanguard_fund_holding":      "vanguard_fund_holding",
			"vanguard_fund_distribution": "vanguard_fund_distribution",
		},
	},
}
