package market

import (
	"github.com/oguzbilgic/fpd"
)

type Account struct {
	// Unique idenifier of the account
	ID string

	// Account balance for each currency
	Currencies map[Currency]*fpd.Decimal

	// Total volume traded during last 30 days
	Volume *fpd.Decimal

	// Account fee % for each trade
	Fee *fpd.Decimal

	Properities []string
}
