package main

import (
	"github.com/goadesign/goa"
	"github.com/rbroggi/go_aggregator/app"
)

// ContractController implements the contract resource.
type ContractController struct {
	*goa.Controller
}

// NewContractController creates a contract controller.
func NewContractController(service *goa.Service) *ContractController {
	return &ContractController{Controller: service.NewController("ContractController")}
}

// Coverage runs the coverage action.
func (c *ContractController) Coverage(ctx *app.CoverageContractContext) error {
	// ContractController_Coverage: start_implement

	// Put your logic here

	res := &app.GoaContractCoverage{}
	return ctx.OKCoverage(res)
	// ContractController_Coverage: end_implement
}

// Exposure runs the exposure action.
func (c *ContractController) Exposure(ctx *app.ExposureContractContext) error {
	// ContractController_Exposure: start_implement

	// Put your logic here

	res := &app.GoaContractExpanded{}
	return ctx.OKExpanded(res)
	// ContractController_Exposure: end_implement
}

// Show runs the show action.
func (c *ContractController) Show(ctx *app.ShowContractContext) error {
	// ContractController_Show: start_implement

	// Put your logic here

	res := &app.GoaTrade{}
	return ctx.OK(res)
	// ContractController_Show: end_implement
}
