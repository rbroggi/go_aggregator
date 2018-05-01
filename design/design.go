package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("exposures", func() { // API defines the microservice endpoint and
	Title("The trade pricing microservice")           // other global properties. There should be one
	Description("A simple goa trade pricing service") // and exactly one API definition appearing in
	Scheme("http")                                    // the design.
	Version("1.0")
	BasePath("/v1")
	Host("localhost:8080")
})

// PriceMedia defines the media type used to render exposures.
var Trade = MediaType("application/vnd.goa.trade+json", func() {
	Description("Trade structure exposures")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("trade_id", String, "Unique trade ID")
		Attribute("contract_id", String, "contract to which trade belong")
		Attribute("counterparty_id", String, "Counterparty to which the trade belongs")
		Attribute("exposures", ArrayOf(Number), "trade exposures")
		Required("trade_id", "counterparty_id", "contract_id", "exposures")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("trade_id") // Media types may have multiple views and must
		Attribute("contract_id")
		Attribute("counterparty_id")
	})
	View("expanded", func() { // View defines a rendering of the media type.
		Attribute("trade_id") // Media types may have multiple views and must
		Attribute("contract_id")
		Attribute("counterparty_id")
		Attribute("exposures")
	})
})

// PriceMedia defines the media type used to render exposures.
var Contract = MediaType("application/vnd.goa.contract+json", func() {
	Description("Trade structure exposures")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("contract_id", String, "Unique contract id")
		Attribute("counterparty_id", String, "Counterparty id to which the contract belongs")
		Attribute("trade_id_list", ArrayOf(String), "List of trades under the contract")
		Attribute("exposures", ArrayOf(Number), "contract aggregated exposures")
		Required("contract_id", "counterparty_id", "trade_id_list", "exposures")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("contract_id") // Media types may have multiple views and must
		Attribute("counterparty_id")
	})
	View("expanded", func() { // View defines a rendering of the media type.
		Attribute("contract_id") // Media types may have multiple views and must
		Attribute("counterparty_id")
		Attribute("trade_id_list")
		Attribute("exposures")
	})
	View("coverage", func() { // View defines a rendering of the media type.
		Attribute("trade_id_list")
	})

})

// PriceMedia defines the media type used to render exposures.
var Counterparty = MediaType("application/vnd.goa.ctp+json", func() {
	Description("Trade structure exposures")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("counterparty_id", String, "Counterparty id")
		Attribute("contract_id_list", ArrayOf(String), "List of contracts under the contract")
		Attribute("exposures", ArrayOf(Number), "counterparty aggregated exposures")
		Required("counterparty_id", "contract_id_list", "exposures")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("counterparty_id") // Media types may have multiple views and must
	})
	View("expanded", func() { // View defines a rendering of the media type.
		Attribute("counterparty_id")
		Attribute("exposures")
	})
	View("coverage", func() { // View defines a rendering of the media type.
		Attribute("contract_id_list")
	})

})

var _ = Resource("trade", func() { // Resources group related API endpoints
	BasePath("/trade")  // together. They map to REST resources for REST
	DefaultMedia(Trade) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get trade by id") // with its path, parameters (both path
		Routing(GET("/show/:tradeID")) // parameters and querystring values) and payload
		Params(func() {                // (shape of the request body).
			Param("tradeID", String, "Trade ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("price", func() {
		Description("Creates new trade and randomly generates its exposures")
		Routing(GET("/price/:ctp/:contract/:trade"))
		Params(func() {
			Param("ctp", String, "Counterparty with which trade is done")
			Param("contract", String, "Contract terms id")
			Param("trade", String, "Trade identifier")
		})
		Response(OK, "text/plain")
	})

	Action("exposure", func() { // Actions define a single API endpoint together
		Description("Get trade exposures by id") // with its path, parameters (both path
		Routing(GET("/exposure/:tradeID"))       // parameters and querystring values) and payload
		Params(func() {                          // (shape of the request body).
			Param("tradeID", String, "Trade ID")
		})
		Response(OK, func() {
			Media(Trade, "expanded")
		}) // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var _ = Resource("contract", func() { // Resources group related API endpoints
	BasePath("/contract") // together. They map to REST resources for REST
	DefaultMedia(Trade)   // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get trade by id")    // with its path, parameters (both path
		Routing(GET("/show/:contractID")) // parameters and querystring values) and payload
		Params(func() {                   // (shape of the request body).
			Param("contractID", String, "Contract ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("exposure", func() { // Actions define a single API endpoint together
		Description("Get contract exposures by id") // with its path, parameters (both path
		Routing(GET("/exposure/:contractID"))       // parameters and querystring values) and payload
		Params(func() {                             // (shape of the request body).
			Param("contractID", String, "Contract ID")
		})
		Response(OK, func() {
			Media(Contract, "expanded")
		}) // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("coverage", func() { // Actions define a single API endpoint together
		Description("Get contract exposures by id") // with its path, parameters (both path
		Routing(GET("/coverage/:contractID"))       // parameters and querystring values) and payload
		Params(func() {                             // (shape of the request body).
			Param("contractID", String, "Contract ID")
		})
		Response(OK, func() {
			Media(Contract, "coverage")
		}) // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})
