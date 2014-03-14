package market

import (
	"github.com/oguzbilgic/fpd"
)

type Broker interface {
	Account() (*Account, error)

	SendOrder(vol *fpd.Decimal, price *Money, orderType OrderType) (string, error)

	CancelOrder(ID string) error

	Orders() ([]*Order, error)
}

type StreamingBroker interface {
	Broker

	NewOrderChan() chan *Order

	NewUserTradeChan() chan *Trade
}
