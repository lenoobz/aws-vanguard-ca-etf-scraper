package ticker

import "fmt"

// GetYahooTicker gets yahoo ticker
func GetYahooTicker(vanguardTicker string) string {
	return fmt.Sprintf("%s.TO", vanguardTicker)
}
