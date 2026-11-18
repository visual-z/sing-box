package include

import (
	"sync"

	"github.com/sagernet/sing-box/adapter/endpoint"
	"github.com/sagernet/sing-box/adapter/inbound"
	"github.com/sagernet/sing-box/adapter/outbound"
)

// Extension hooks let a downstream build register protocol types that live
// outside this module. Every registry constructor in this package builds a
// fresh registry (libbox does so for each service and for CheckConfig), so a
// downstream package cannot add types after the fact; it registers a callback
// here, usually from init, and every new registry runs it.
//
// Carried by the Parallux fork (github.com/visual-z/sing-box, branch
// parallux/<version>) as its only change.

var (
	extensionAccess   sync.Mutex
	inboundExtension  []func(*inbound.Registry)
	outboundExtension []func(*outbound.Registry)
	endpointExtension []func(*endpoint.Registry)
)

// RegisterInboundExtension adds inbound types to every registry built by
// InboundRegistry from now on.
func RegisterInboundExtension(register func(*inbound.Registry)) {
	extensionAccess.Lock()
	defer extensionAccess.Unlock()
	inboundExtension = append(inboundExtension, register)
}

// RegisterOutboundExtension adds outbound types to every registry built by
// OutboundRegistry from now on.
func RegisterOutboundExtension(register func(*outbound.Registry)) {
	extensionAccess.Lock()
	defer extensionAccess.Unlock()
	outboundExtension = append(outboundExtension, register)
}

// RegisterEndpointExtension adds endpoint types to every registry built by
// EndpointRegistry from now on.
func RegisterEndpointExtension(register func(*endpoint.Registry)) {
	extensionAccess.Lock()
	defer extensionAccess.Unlock()
	endpointExtension = append(endpointExtension, register)
}

func applyExtensions[R any](registry R, extensions *[]func(R)) {
	extensionAccess.Lock()
	registrations := append([]func(R){}, *extensions...)
	extensionAccess.Unlock()
	for _, register := range registrations {
		register(registry)
	}
}
