package consts

import "fmt"

// SchemaVersion const
const SchemaVersion = 1

// CollectionFundOverview const
const CollectionFundOverview = "fund_overview"

// TimeoutMS const
const TimeoutMS = 15000

// MinPoolSize const
const MinPoolSize = 5

// MaxPoolSize const
const MaxPoolSize = 10

// MaxIdleTimeMS const
const MaxIdleTimeMS = 360000

// Username const
const Username = "lenoob_dev"

// Password const
const Password = "lenoob_dev"

// Host const
const Host = "lenoobetfdevcluster.jd7wd.mongodb.net"

// Dbname const
const Dbname = "etf_funds"

// AllowDomain const
const AllowDomain = "api.vanguard.com"

// DomainGlob const
const DomainGlob = "*vanguard.*"

// FundListURL const
const FundListURL = "https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-listview-data-en.json"

// GetFundOverviewURL get fund overview url
func GetFundOverviewURL(portID string) string {
	return fmt.Sprintf("https://api.vanguard.com/rs/gre/gra/1.7.0/datasets/caw-indv-overview-data-etf.json?vars=portId:%s,lang:en&path=[portId=%s][0]", portID, portID)
}
