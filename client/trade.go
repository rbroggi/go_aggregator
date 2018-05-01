// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "exposures": trade Resource Client
//
// Command:
// $ goagen
// --design=github.com/rbroggi/go_aggregator/design
// --out=$(GOPATH)/src/github.com/rbroggi/go_aggregator
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ExposureTradePath computes a request path to the exposure action of trade.
func ExposureTradePath(tradeID string) string {
	param0 := tradeID

	return fmt.Sprintf("/v1/trade/exposure/%s", param0)
}

// Get trade exposures by id
func (c *Client) ExposureTrade(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewExposureTradeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewExposureTradeRequest create the request corresponding to the exposure action endpoint of the trade resource.
func (c *Client) NewExposureTradeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PriceTradePath computes a request path to the price action of trade.
func PriceTradePath(ctp string, contract string, trade string) string {
	param0 := ctp
	param1 := contract
	param2 := trade

	return fmt.Sprintf("/v1/trade/price/%s/%s/%s", param0, param1, param2)
}

// Creates new trade and randomly generates its exposures
func (c *Client) PriceTrade(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPriceTradeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPriceTradeRequest create the request corresponding to the price action endpoint of the trade resource.
func (c *Client) NewPriceTradeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowTradePath computes a request path to the show action of trade.
func ShowTradePath(tradeID string) string {
	param0 := tradeID

	return fmt.Sprintf("/v1/trade/show/%s", param0)
}

// Get trade by id
func (c *Client) ShowTrade(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTradeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTradeRequest create the request corresponding to the show action endpoint of the trade resource.
func (c *Client) NewShowTradeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}