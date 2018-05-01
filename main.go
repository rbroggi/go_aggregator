//go:generate goagen bootstrap -d github.com/rbroggi/go_aggregator/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/rbroggi/go_aggregator/app"
)

func main() {
	// Create service
	service := goa.New("exposures")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "trade" controller
	c := NewTradeController(service)
	app.MountTradeController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
