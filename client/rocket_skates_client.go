package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/client/bootenvs"
	"github.com/rackn/rocket-skates/client/dhcp_leases"
	"github.com/rackn/rocket-skates/client/files"
	"github.com/rackn/rocket-skates/client/isos"
	"github.com/rackn/rocket-skates/client/machines"
	"github.com/rackn/rocket-skates/client/templates"
)

// Default rocket skates HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new rocket skates HTTP client.
func NewHTTPClient(formats strfmt.Registry) *RocketSkates {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("127.0.0.1:8092", "/api/v3", []string{"https"})
	return New(transport, formats)
}

// New creates a new rocket skates client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *RocketSkates {
	cli := new(RocketSkates)
	cli.Transport = transport

	cli.Bootenvs = bootenvs.New(transport, formats)

	cli.DhcpLeases = dhcp_leases.New(transport, formats)

	cli.Files = files.New(transport, formats)

	cli.Isos = isos.New(transport, formats)

	cli.Machines = machines.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

	return cli
}

// RocketSkates is a client for rocket skates
type RocketSkates struct {
	Bootenvs *bootenvs.Client

	DhcpLeases *dhcp_leases.Client

	Files *files.Client

	Isos *isos.Client

	Machines *machines.Client

	Templates *templates.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *RocketSkates) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Bootenvs.SetTransport(transport)

	c.DhcpLeases.SetTransport(transport)

	c.Files.SetTransport(transport)

	c.Isos.SetTransport(transport)

	c.Machines.SetTransport(transport)

	c.Templates.SetTransport(transport)

}