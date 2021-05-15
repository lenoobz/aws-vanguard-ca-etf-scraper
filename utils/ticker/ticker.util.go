package ticker

import "fmt"

// GenYahooTickerFromVanguardTicker gen yahoo ticker from vanguard ticker
func GenYahooTickerFromVanguardTicker(vanguardTicker string) string {
	return fmt.Sprintf("%s.TO", vanguardTicker)
}
