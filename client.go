package osincli

import (
	"net/http"
)

type Client struct {
	// Client configuration
	Config *ClientConfig

	// Transport is the HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// Creates a new client
func NewClient(config *ClientConfig) *Client {
	return &Client{
		Config: config,
	}
}
