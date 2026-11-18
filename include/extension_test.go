package include

import (
	"context"
	"testing"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/adapter/outbound"
	"github.com/sagernet/sing-box/log"
)

type extensionTestOptions struct{}

func TestOutboundRegistryRunsRegisteredExtensions(t *testing.T) {
	const outboundType = "extension-test"
	if _, loaded := OutboundRegistry().CreateOptions(outboundType); loaded {
		t.Fatalf("%s is registered before the extension is added", outboundType)
	}
	RegisterOutboundExtension(func(registry *outbound.Registry) {
		outbound.Register[extensionTestOptions](registry, outboundType, func(context.Context, adapter.Router, log.ContextLogger, string, extensionTestOptions) (adapter.Outbound, error) {
			return nil, nil
		})
	})
	if _, loaded := OutboundRegistry().CreateOptions(outboundType); !loaded {
		t.Fatalf("a registry built after RegisterOutboundExtension does not know %s", outboundType)
	}
}
