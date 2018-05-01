package main

import (
	"github.com/goadesign/goa"
	"github.com/rbroggi/go_aggregator/app"
	"math/rand"
)

const (
	bound = 100
)

var tradeRegistry = make(map[string]app.GoaTradeExpanded)

// TradeController implements the trade resource.
type TradeController struct {
	*goa.Controller
}

// NewTradeController creates a trade controller.
func NewTradeController(service *goa.Service) *TradeController {
	return &TradeController{Controller: service.NewController("TradeController")}
}

// Exposure runs the exposure action.
func (c *TradeController) Exposure(ctx *app.ExposureTradeContext) error {
	// TradeController_Exposure: start_implement

	t, ok := tradeRegistry[ctx.TradeID]
	if !ok {
		return ctx.NotFound()
	}

	res := &app.GoaTradeExpanded{
		TradeID:        t.TradeID,
		ContractID:     t.ContractID,
		CounterpartyID: t.CounterpartyID,
		Exposures:      t.Exposures,
	}
	return ctx.OKExpanded(res)
	// TradeController_Exposure: end_implement
}

// Price runs the price action.
func (c *TradeController) Price(ctx *app.PriceTradeContext) error {
	// TradeController_Price: start_implement
	tradeRegistry[ctx.Trade] = app.GoaTradeExpanded{
		TradeID:        ctx.Trade,
		ContractID:     ctx.Contract,
		CounterpartyID: ctx.Ctp,
		Exposures:      price(),
	}

	// Put your logic here

	return nil
	// TradeController_Price: end_implement
}

// Show runs the show action.
func (c *TradeController) Show(ctx *app.ShowTradeContext) error {
	// TradeController_Show: start_implement

	// Put your logic here
	t, ok := tradeRegistry[ctx.TradeID]
	if !ok {
		return ctx.NotFound()
	}
	res := &app.GoaTrade{
		TradeID:        t.TradeID,
		ContractID:     t.ContractID,
		CounterpartyID: t.CounterpartyID,
	}
	return ctx.OK(res)
	// TradeController_Show: end_implement
}

func price() []float64 {
	slice := make([]float64, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, -bound+2*bound*rand.Float64())
	}
	return slice
}
