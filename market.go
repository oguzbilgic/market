package market

type Market interface {
	Ticker() (*Tick, error)

	OrderBook() ([]*Depth, error)
}

type StreamingMarket interface {
	Market

	NewTickChan() chan *Tick

	NewDepthChan() chan *Depth

	NewTradeChan() chan *Trade
}
