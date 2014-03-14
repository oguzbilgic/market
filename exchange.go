package market

type Exchange interface {
	Broker

	Market
}

type StreamingExchange interface {
	StreamingBroker

	StreamingMarket
}
