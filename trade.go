package market

import (
	"github.com/oguzbilgic/fpd"
	"time"
)

type Trade struct {
	// Unique idenifier of the trade
	ID string

	// Symbol of the security traded
	Symbol string

	// Currency used for the trade
	Currency Currency

	// Volume of the trade
	Volume *fpd.Decimal

	// Price at which trade occured
	Price *fpd.Decimal

	// Time when the trade occured
	Time time.Time

	// Type of the order caused the trade
	Type OrderType
}
