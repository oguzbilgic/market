package market

import (
	"github.com/oguzbilgic/fpd"
)

type Money struct {
	Amount *fpd.Decimal

	Currency Currency
}

func (m *Money) String() string {
	return m.Amount.FormattedString() + " " + string(m.Currency)
}
