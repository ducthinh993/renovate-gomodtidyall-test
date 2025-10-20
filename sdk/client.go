package sdk

import (
	"github.com/ducthinh993/renovate-gomodtidyall-test/shared"
)

// Client provides SDK functionality using shared library
type Client struct {
	name string
}

// NewClient creates a new SDK client
func NewClient(name string) *Client {
	return &Client{name: name}
}

// ProcessMessage uses shared library to format messages
func (c *Client) ProcessMessage(msg string) shared.Response {
	return shared.NewResponse("processed", "SDK: "+msg)
}

// GetName returns the client name
func (c *Client) GetName() string {
	return c.name
}