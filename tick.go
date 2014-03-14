package market

import (
	"github.com/oguzbilgic/fpd"
	"time"
)

type Tick struct {
	// Symbol of the security
	Symbol string

	// Currency of the prices
	Currency Currency

	//
	Time time.Time

	// Market volume
	Volume *fpd.Decimal

	High *fpd.Decimal

	Low *fpd.Decimal

	// Highest bid price
	Bid *fpd.Decimal

	// Lowest ask price
	Ask *fpd.Decimal

	// Last traded price
	Last *fpd.Decimal
}
