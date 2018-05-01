package aggregator

import "github.com/rbroggi/go_aggregator/app"

type Aggregator interface {
	Aggregate(t app.GoaTradeExpanded)
}

func NewAggregator() Aggregator {
	return &AGE{
		ctpRegistry:      make(map[string]app.GoaCtpExpanded),
		contractRegistry: make(map[string]app.GoaContractExpanded),
		tradeRegistry:    make(map[string]app.GoaTradeExpanded),
	}

}

type AGE struct {
	ctpRegistry      map[string]app.GoaCtpExpanded
	contractRegistry map[string]app.GoaContractExpanded
	tradeRegistry    map[string]app.GoaTradeExpanded
}

func (a *AGE) Aggregate(t app.GoaTradeExpanded) {
	tp, ok := a.tradeRegistry[t.TradeID]
	offset := make([]float64, 10)
	if ok {
		offset = IndexWiseFunc(tp.Exposures, t.Exposures, func(a, b float64) float64 { return a - b })
	} else {
		offset = t.Exposures
	}

	con, ok := a.contractRegistry[t.ContractID]
	if ok {
		con.Exposures = IndexWiseFunc(con.Exposures, offset, func(a, b float64) float64 { return a + b })
	} else {
		con.Exposures = offset
	}

	ctp, ok := a.contractRegistry[t.CounterpartyID]
	if ok {
		ctp.Exposures = IndexWiseFunc(ctp.Exposures, offset, func(a, b float64) float64 { return a + b })
	} else {
		ctp.Exposures = offset
	}

}

//perform index-wise operations determined by f
func IndexWiseFunc(a, b []float64, f func(j, k float64) float64) []float64 {
	c := make([]float64, len(a))
	for i, _ := range a {
		c[i] = f(a[i], b[i])
	}
	return c
}
