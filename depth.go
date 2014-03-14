package market

import (
	"github.com/oguzbilgic/fpd"
)

type Depth struct {
	// Symbol of the security
	Symbol string

	// Currency of the depth
	Currency Currency

	// Change in volume for the given price depth
	// Volume is positive when new order(s) is added
	// Volume is negative when the order(s) is canceled
	Volume *fpd.Decimal

	// Price of the depth
	Price *fpd.Decimal

	// Type of the depth
	Type OrderType
}
