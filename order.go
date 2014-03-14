package market

import (
	"github.com/oguzbilgic/fpd"
	"time"
)

type OrderType string

const (
	ASK OrderType = "ask"
	BID           = "bid"
)

type OrderStatus string

const (
	PENDING   OrderStatus = "pending"
	OPEN                  = "open"
	EXECUTING             = "executing"
	INVALID               = "invalid"
	STOP                  = "stop"
)

type Order struct {
	// Unique idenifier of the order
	ID string

	// Symbol of the security ordered
	Symbol string

	// Currency used for the order
	Currency Currency

	// Volume of the order
	Volume *fpd.Decimal

	// Price of the order
	Price *fpd.Decimal

	// Time when the order is received by the market
	Time time.Time

	// Type of the order
	Type OrderType

	// Status of the order
	Status OrderStatus
}
